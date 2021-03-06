// Code generated by "stringer -type=ConnectionStatus ./api/peer"; DO NOT EDIT.

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

package peer

import "fmt"

const _ConnectionStatus_name = "UnavailableConnectingAvailable"

var _ConnectionStatus_index = [...]uint8{0, 11, 21, 30}

func (i ConnectionStatus) String() string {
	if i < 0 || i >= ConnectionStatus(len(_ConnectionStatus_index)-1) {
		return fmt.Sprintf("ConnectionStatus(%d)", i)
	}
	return _ConnectionStatus_name[_ConnectionStatus_index[i]:_ConnectionStatus_index[i+1]]
}
