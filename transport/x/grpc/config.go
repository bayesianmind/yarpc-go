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

package grpc

import (
	"fmt"
	"net"

	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/x/config"
)

const transportName = "grpc"

// TransportSpec returns a TransportSpec for the gRPC transport.
//
// See TransportConfig, InboundConfig, and OutboundConfig for details on the
// different configuration parameters supported by this Transport.
//
// Any TransportOption, InboundOption, or OutboundOption may be passed to this function.
// These options will be applied BEFORE configuration parameters are
// interpreted. This allows configuration parameters to override Options
// provided to TransportSpec.
func TransportSpec(opts ...Option) config.TransportSpec {
	transportSpec, err := newTransportSpec(opts...)
	if err != nil {
		panic(err.Error())
	}
	return config.TransportSpec{
		Name:               transportName,
		BuildTransport:     transportSpec.buildTransport,
		BuildInbound:       transportSpec.buildInbound,
		BuildUnaryOutbound: transportSpec.buildUnaryOutbound,
	}
}

// TransportConfig configures a gRPC Transport.
//
// This is currently just a placeholder.
type TransportConfig struct{}

// InboundConfig configures a gRPC Inbound.
//
// inbounds:
//   grpc:
//     address: ":80"
type InboundConfig struct {
	// Address to listen on. This field is required.
	Address string `config:"address,interpolate"`
}

// OutboundConfig configures a gRPC Outbound.
//
// outbounds:
//   myservice:
//     grpc:
//       address: ":80"
type OutboundConfig struct {
	// Address to connect to. This field is required.
	Address string `config:"address,interpolate"`
}

type transportSpec struct {
	TransportOptions []TransportOption
	InboundOptions   []InboundOption
	OutboundOptions  []OutboundOption
}

func newTransportSpec(opts ...Option) (*transportSpec, error) {
	transportSpec := &transportSpec{}
	for _, o := range opts {
		switch opt := o.(type) {
		case TransportOption:
			transportSpec.TransportOptions = append(transportSpec.TransportOptions, opt)
		case InboundOption:
			transportSpec.InboundOptions = append(transportSpec.InboundOptions, opt)
		case OutboundOption:
			transportSpec.OutboundOptions = append(transportSpec.OutboundOptions, opt)
		default:
			return nil, fmt.Errorf("unknown option of type %T: %v", o, o)
		}
	}
	return transportSpec, nil
}

func (t *transportSpec) buildTransport(*TransportConfig, *config.Kit) (transport.Transport, error) {
	return NewTransport(t.TransportOptions...), nil
}

func (t *transportSpec) buildInbound(inboundConfig *InboundConfig, tr transport.Transport, _ *config.Kit) (transport.Inbound, error) {
	trans, ok := tr.(*Transport)
	if !ok {
		return nil, newTransportCastError(tr)
	}
	if inboundConfig.Address == "" {
		return nil, newRequiredFieldMissingError("address")
	}
	listener, err := net.Listen("tcp", inboundConfig.Address)
	if err != nil {
		return nil, err
	}
	return trans.NewInbound(listener, t.InboundOptions...), nil
}

func (t *transportSpec) buildUnaryOutbound(outboundConfig *OutboundConfig, tr transport.Transport, _ *config.Kit) (transport.UnaryOutbound, error) {
	trans, ok := tr.(*Transport)
	if !ok {
		return nil, newTransportCastError(tr)
	}
	if outboundConfig.Address == "" {
		return nil, newRequiredFieldMissingError("address")
	}
	return trans.NewSingleOutbound(outboundConfig.Address, t.OutboundOptions...), nil
}

func newTransportCastError(tr transport.Transport) error {
	return fmt.Errorf("could not cast %T to a *grpc.Transport", tr)
}

func newRequiredFieldMissingError(field string) error {
	return fmt.Errorf("required field missing: %s", field)
}
