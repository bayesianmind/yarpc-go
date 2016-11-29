// Code generated by thriftrw v1.0.0
// @generated

// Copyright (c) 2016 Uber Technologies, Inc.
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

type ThriftTest_TestMultiException_Args struct {
	Arg0 *string `json:"arg0,omitempty"`
	Arg1 *string `json:"arg1,omitempty"`
}

func (v *ThriftTest_TestMultiException_Args) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg0 != nil {
		w, err = wire.NewValueString(*(v.Arg0)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Arg1 != nil {
		w, err = wire.NewValueString(*(v.Arg1)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ThriftTest_TestMultiException_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Arg0 = &x
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Arg1 = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *ThriftTest_TestMultiException_Args) String() string {
	var fields [2]string
	i := 0
	if v.Arg0 != nil {
		fields[i] = fmt.Sprintf("Arg0: %v", *(v.Arg0))
		i++
	}
	if v.Arg1 != nil {
		fields[i] = fmt.Sprintf("Arg1: %v", *(v.Arg1))
		i++
	}
	return fmt.Sprintf("ThriftTest_TestMultiException_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestMultiException_Args) MethodName() string {
	return "testMultiException"
}

func (v *ThriftTest_TestMultiException_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var ThriftTest_TestMultiException_Helper = struct {
	Args           func(arg0 *string, arg1 *string) *ThriftTest_TestMultiException_Args
	IsException    func(error) bool
	WrapResponse   func(*Xtruct, error) (*ThriftTest_TestMultiException_Result, error)
	UnwrapResponse func(*ThriftTest_TestMultiException_Result) (*Xtruct, error)
}{}

func init() {
	ThriftTest_TestMultiException_Helper.Args = func(arg0 *string, arg1 *string) *ThriftTest_TestMultiException_Args {
		return &ThriftTest_TestMultiException_Args{Arg0: arg0, Arg1: arg1}
	}
	ThriftTest_TestMultiException_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *Xception:
			return true
		case *Xception2:
			return true
		default:
			return false
		}
	}
	ThriftTest_TestMultiException_Helper.WrapResponse = func(success *Xtruct, err error) (*ThriftTest_TestMultiException_Result, error) {
		if err == nil {
			return &ThriftTest_TestMultiException_Result{Success: success}, nil
		}
		switch e := err.(type) {
		case *Xception:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for ThriftTest_TestMultiException_Result.Err1")
			}
			return &ThriftTest_TestMultiException_Result{Err1: e}, nil
		case *Xception2:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for ThriftTest_TestMultiException_Result.Err2")
			}
			return &ThriftTest_TestMultiException_Result{Err2: e}, nil
		}
		return nil, err
	}
	ThriftTest_TestMultiException_Helper.UnwrapResponse = func(result *ThriftTest_TestMultiException_Result) (success *Xtruct, err error) {
		if result.Err1 != nil {
			err = result.Err1
			return
		}
		if result.Err2 != nil {
			err = result.Err2
			return
		}
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type ThriftTest_TestMultiException_Result struct {
	Success *Xtruct    `json:"success,omitempty"`
	Err1    *Xception  `json:"err1,omitempty"`
	Err2    *Xception2 `json:"err2,omitempty"`
}

func (v *ThriftTest_TestMultiException_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
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
	if v.Err1 != nil {
		w, err = v.Err1.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.Err2 != nil {
		w, err = v.Err2.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestMultiException_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Xception2_Read(w wire.Value) (*Xception2, error) {
	var v Xception2
	err := v.FromWire(w)
	return &v, err
}

func (v *ThriftTest_TestMultiException_Result) FromWire(w wire.Value) error {
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
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Err1, err = _Xception_Read(field.Value)
				if err != nil {
					return err
				}
			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.Err2, err = _Xception2_Read(field.Value)
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
	if v.Err1 != nil {
		count++
	}
	if v.Err2 != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ThriftTest_TestMultiException_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *ThriftTest_TestMultiException_Result) String() string {
	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	if v.Err1 != nil {
		fields[i] = fmt.Sprintf("Err1: %v", v.Err1)
		i++
	}
	if v.Err2 != nil {
		fields[i] = fmt.Sprintf("Err2: %v", v.Err2)
		i++
	}
	return fmt.Sprintf("ThriftTest_TestMultiException_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestMultiException_Result) MethodName() string {
	return "testMultiException"
}

func (v *ThriftTest_TestMultiException_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}