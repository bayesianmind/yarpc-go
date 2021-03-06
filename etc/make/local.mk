# Paths besides auto-detected generated files that should be excluded from
# lint results.
LINT_EXCLUDES_EXTRAS =

# Regex for 'go vet' rules to ignore
GOVET_IGNORE_RULES = \
	possible formatting directive in Error call

ERRCHECK_FLAGS := -ignoretests
ERRCHECK_EXCLUDES := \.Close\(\) \.Stop\(\)
FILTER_ERRCHECK := grep -v $(patsubst %,-e %, $(ERRCHECK_EXCLUDES))

# The number of jobs allocated to run examples tests in parallel
# The goal is to have all examples tests run in parallel, and
# this is currently greater than the number of examples tests
EXAMPLES_JOBS ?= 16

GEN_BINS_INTERNAL = $(BIN)/thriftrw-plugin-yarpc $(BIN)/protoc-gen-yarpc-go

$(BIN)/thriftrw-plugin-yarpc: ./encoding/thrift/thriftrw-plugin-yarpc/*.go
	@mkdir -p $(BIN)
	go build -o $(BIN)/thriftrw-plugin-yarpc ./encoding/thrift/thriftrw-plugin-yarpc

$(BIN)/protoc-gen-yarpc-go: ./encoding/protobuf/protoc-gen-yarpc-go/*.go
	@mkdir -p $(BIN)
	go build -o $(BIN)/protoc-gen-yarpc-go ./encoding/protobuf/protoc-gen-yarpc-go

.PHONY: build
build: __eval_packages ## go build all packages
	go build $(PACKAGES)

.PHONY: generate
generate: $(GEN_BINS) $(GEN_BINS_INTERNAL) ## call generation script
	@go get github.com/golang/mock/mockgen
	@PATH=$(BIN):$$PATH ./etc/bin/generate.sh
ifdef WITHIN_DOCKER
	@chown -R --reference . .
endif

.PHONY: nogogenerate
nogogenerate: __eval_go_files ## check to make sure go:generate is not used
	$(eval NOGOGENERATE_LOG := $(shell mktemp -t nogogenerate.XXXXX))
	@grep -n \/\/go:generate $(GO_FILES) 2>&1 > $(NOGOGENERATE_LOG) || true
	@[ ! -s "$(NOGOGENERATE_LOG)" ] || (echo "do not use //go:generate, add to etc/bin/generate.sh instead:" | cat - $(NOGOGENERATE_LOG) && false)

.PHONY: generatenodiff
generatenodiff: ## make sure no diff is generated by make generate
	$(eval GENERATENODIFF_PRE := $(shell mktemp -t generatenodiff_pre.XXXXX))
	$(eval GENERATENODIFF_POST := $(shell mktemp -t generatenodiff_post.XXXXX))
	$(eval GENERATENODIFF_DIFF := $(shell mktemp -t generatenodiff_diff.XXXXX))
	@git status --short > $(GENERATENODIFF_PRE)
	@$(MAKE) generate
	@git status --short > $(GENERATENODIFF_POST)
	@diff $(GENERATENODIFF_PRE) $(GENERATENODIFF_POST) > $(GENERATENODIFF_DIFF) || true
	@[ ! -s "$(GENERATENODIFF_DIFF)" ] || (echo "make generate produced a diff, make sure to check these in:" | cat - $(GENERATENODIFF_DIFF) && false)


.PHONY: gofmt
gofmt: __eval_go_files ## check gofmt
	$(eval FMT_LOG := $(shell mktemp -t gofmt.XXXXX))
	@PATH=$(BIN):$$PATH gofmt -e -s -l $(GO_FILES) | $(FILTER_LINT) > $(FMT_LOG) || true
	@[ ! -s "$(FMT_LOG)" ] || (echo "gofmt failed:" | cat - $(FMT_LOG) && false)

.PHONY: govet
govet: __eval_packages __eval_go_files ## check go vet
	$(eval VET_LOG := $(shell mktemp -t govet.XXXXX))
	@go vet $(PACKAGES) 2>&1 \
		| grep -v '^exit status' \
		| grep -v "$(GOVET_IGNORE_RULES)" \
		| $(FILTER_LINT) > $(VET_LOG) || true
	@[ ! -s "$(VET_LOG)" ] || (echo "govet failed:" | cat - $(VET_LOG) && false)

.PHONY: golint
golint: $(GOLINT) __eval_packages __eval_go_files ## check golint
	$(eval LINT_LOG := $(shell mktemp -t golint.XXXXX))
	@for pkg in $(PACKAGES); do \
		PATH=$(BIN):$$PATH golint $$pkg | $(FILTER_LINT) >> $(LINT_LOG) || true; \
	done
	@[ ! -s "$(LINT_LOG)" ] || (echo "golint failed:" | cat - $(LINT_LOG) && false)

.PHONY: staticcheck
staticcheck: $(STATICCHECK) __eval_packages __eval_go_files ## check staticcheck
	$(eval STATICCHECK_LOG := $(shell mktemp -t staticcheck.XXXXX))
	@PATH=$(BIN):$$PATH staticcheck $(PACKAGES) 2>&1 | $(FILTER_LINT) > $(STATICCHECK_LOG) || true
	@[ ! -s "$(STATICCHECK_LOG)" ] || (echo "staticcheck failed:" | cat - $(STATICCHECK_LOG) && false)

.PHONY: errcheck
errcheck: $(ERRCHECK) __eval_packages __eval_go_files ## check errcheck
	$(eval ERRCHECK_LOG := $(shell mktemp -t errcheck.XXXXX))
	@PATH=$(BIN):$$PATH errcheck $(ERRCHECK_FLAGS) $(PACKAGES) 2>&1 | $(FILTER_LINT) | $(FILTER_ERRCHECK) > $(ERRCHECK_LOG) || true
	@[ ! -s "$(ERRCHECK_LOG)" ] || (echo "errcheck failed:" | cat - $(ERRCHECK_LOG) && false)

.PHONY: verifyversion
verifyversion: ## verify the version in the changelog is the same as in version.go
	$(eval CHANGELOG_VERSION := $(shell grep '^v[0-9]' CHANGELOG.md | head -n1 | cut -d' ' -f1))
	$(eval INTHECODE_VERSION := $(shell perl -ne '/^const Version.*"([^"]+)".*$$/ && print "v$$1\n"' version.go))
	@if [ "$(INTHECODE_VERSION)" = "$(CHANGELOG_VERSION)" ]; then \
		echo "yarpc-go: $(CHANGELOG_VERSION)"; \
	elif [ "$(INTHECODE_VERSION)" = "$(CHANGELOG_VERSION)-dev" ]; then \
		echo "yarpc-go (development): $(INTHECODE_VERSION)"; \
	else \
		echo "Version number in version.go does not match CHANGELOG.md"; \
		echo "version.go: $(INTHECODE_VERSION)"; \
		echo "CHANGELOG : $(CHANGELOG_VERSION)"; \
		exit 1; \
	fi

.PHONY: verifycodecovignores
verifycodecovignores: ## verify that .codecov.yml contains all .nocover packages
	@find . '(' -name vendor -o -name .glide ')' -prune -o -name .nocover -exec dirname '{}' ';' \
		| cut -b2- \
		| while read f; do \
			if ! grep "$$f" .codecov.yml >/dev/null; then \
				echo ".codecov.yml is out of date: add $$f to it"; \
				exit 1; \
			fi \
		done

.PHONY: basiclint
basiclint: gofmt govet golint staticcheck errcheck # run gofmt govet golint staticcheck errcheck

.PHONY: lint
lint: basiclint generatenodiff nogogenerate verifyversion verifycodecovignores ## run all linters

.PHONY: test
test: $(THRIFTRW) __eval_packages ## run all tests
	PATH=$(BIN):$$PATH go test -race $(PACKAGES)

.PHONY: cover
cover: $(THRIFTRW) $(GOCOVMERGE) $(PARALLEL_EXEC) $(COVER) __eval_packages ## run all tests and output code coverage
	PATH=$(BIN):$$PATH ./etc/bin/cover.sh $(PACKAGES)
	go tool cover -html=coverage.main.txt -o cover.main.html
	go tool cover -html=coverage.x.txt -o cover.x.html

.PHONY: codecov
codecov: SHELL := /bin/bash
codecov: cover ## run code coverage and upload to codecov.io
	bash <(curl -s https://codecov.io/bash) -c -f coverage.main.txt -F main
	bash <(curl -s https://codecov.io/bash) -c -f coverage.x.txt -F experimental

.PHONY: examples
examples: ## run all examples tests
	RUN=$(RUN) V=$(V) $(MAKE) -j $(EXAMPLES_JOBS) -C internal/examples

.PHONY: __eval_packages
__eval_packages:
ifndef PACKAGES
	$(eval PACKAGES := $(shell go list ./... | grep -v go\.uber\.org\/yarpc\/vendor))
endif

.PHONY: __eval_go_files
__eval_go_files:
	$(eval GO_FILES := $(shell find . -name '*.go' | sed 's/^\.\///' | grep -v -e ^vendor\/ -e ^\.glide\/))
	$(eval GENERATED_GO_FILES := $(shell \
		find $(GO_FILES) \
		-exec sh -c 'head -n30 {} | grep "Code generated by\|Autogenerated by\|Automatically generated by\|@generated" >/dev/null' \; \
		-print))
	$(eval FILTER_LINT := grep -v $(patsubst %,-e %, $(GENERATED_GO_FILES) $(LINT_EXCLUDES_EXTRAS)))
