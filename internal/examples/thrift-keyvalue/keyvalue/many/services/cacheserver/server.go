// Code generated by thriftrw-plugin-yarpc
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

package cacheserver

import (
	"context"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/internal/examples/thrift-keyvalue/keyvalue/many/services"
	"go.uber.org/yarpc/encoding/thrift"
)

// Interface is the server-side interface for the Cache service.
type Interface interface {
	Clear(
		ctx context.Context,
	) error

	ClearAfter(
		ctx context.Context,
		DurationMS *int64,
	) error
}

// New prepares an implementation of the Cache service for
// registration.
//
// 	handler := CacheHandler{}
// 	dispatcher.Register(cacheserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "Cache",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "clear",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Oneway,
					Oneway: thrift.OnewayHandler(h.Clear),
				},
				Signature:    "Clear()",
				ThriftModule: services.ThriftModule,
			},

			thrift.Method{
				Name: "clearAfter",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Oneway,
					Oneway: thrift.OnewayHandler(h.ClearAfter),
				},
				Signature:    "ClearAfter(DurationMS *int64)",
				ThriftModule: services.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 2)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

func (h handler) Clear(ctx context.Context, body wire.Value) error {
	var args services.Cache_Clear_Args
	if err := args.FromWire(body); err != nil {
		return err
	}

	return h.impl.Clear(ctx)
}

func (h handler) ClearAfter(ctx context.Context, body wire.Value) error {
	var args services.Cache_ClearAfter_Args
	if err := args.FromWire(body); err != nil {
		return err
	}

	return h.impl.ClearAfter(ctx, args.DurationMS)
}
