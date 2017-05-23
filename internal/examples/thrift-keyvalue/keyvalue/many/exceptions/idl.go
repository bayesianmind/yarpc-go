// Code generated by thriftrw v1.2.0
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package exceptions

import "go.uber.org/thriftrw/thriftreflect"

var ThriftModule = &thriftreflect.ThriftModule{Name: "exceptions", Package: "go.uber.org/yarpc/internal/examples/thrift-keyvalue/keyvalue/many/exceptions", FilePath: "many/exceptions.thrift", SHA1: "7515e1a7bcd9ef547ad37a0eaaeaf89df21897fa", Raw: rawIDL}

const rawIDL = "exception EmptyException {}\n\nexception DoesNotExistException {\n    1: required string key\n    2: optional string Error (go.name=\"Error2\")\n}\n"
