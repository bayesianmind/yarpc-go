// Code generated by thriftrw v1.3.0
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

package gauntlet

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ThriftTest_TestStruct_Args struct {
	Thing *Xtruct `json:"thing,omitempty"`
}

func (v *ThriftTest_TestStruct_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Thing != nil {
		w, err = v.Thing.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ThriftTest_TestStruct_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Thing, err = _Xtruct_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *ThriftTest_TestStruct_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Thing != nil {
		fields[i] = fmt.Sprintf("Thing: %v", v.Thing)
		i++
	}
	return fmt.Sprintf("ThriftTest_TestStruct_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestStruct_Args) Equals(rhs *ThriftTest_TestStruct_Args) bool {
	if !((v.Thing == nil && rhs.Thing == nil) || (v.Thing != nil && rhs.Thing != nil && v.Thing.Equals(rhs.Thing))) {
		return false
	}
	return true
}

func (v *ThriftTest_TestStruct_Args) MethodName() string {
	return "testStruct"
}

func (v *ThriftTest_TestStruct_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var ThriftTest_TestStruct_Helper = struct {
	Args           func(thing *Xtruct) *ThriftTest_TestStruct_Args
	IsException    func(error) bool
	WrapResponse   func(*Xtruct, error) (*ThriftTest_TestStruct_Result, error)
	UnwrapResponse func(*ThriftTest_TestStruct_Result) (*Xtruct, error)
}{}

func init() {
	ThriftTest_TestStruct_Helper.Args = func(thing *Xtruct) *ThriftTest_TestStruct_Args {
		return &ThriftTest_TestStruct_Args{Thing: thing}
	}
	ThriftTest_TestStruct_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	ThriftTest_TestStruct_Helper.WrapResponse = func(success *Xtruct, err error) (*ThriftTest_TestStruct_Result, error) {
		if err == nil {
			return &ThriftTest_TestStruct_Result{Success: success}, nil
		}
		return nil, err
	}
	ThriftTest_TestStruct_Helper.UnwrapResponse = func(result *ThriftTest_TestStruct_Result) (success *Xtruct, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type ThriftTest_TestStruct_Result struct {
	Success *Xtruct `json:"success,omitempty"`
}

func (v *ThriftTest_TestStruct_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestStruct_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ThriftTest_TestStruct_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _Xtruct_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ThriftTest_TestStruct_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *ThriftTest_TestStruct_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("ThriftTest_TestStruct_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestStruct_Result) Equals(rhs *ThriftTest_TestStruct_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	return true
}

func (v *ThriftTest_TestStruct_Result) MethodName() string {
	return "testStruct"
}

func (v *ThriftTest_TestStruct_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
