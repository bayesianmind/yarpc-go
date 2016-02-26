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

package raw

import (
	"bytes"
	"io/ioutil"

	"github.com/yarpc/yarpc-go/transport"
)

// Client makes Raw requests to a single service.
type Client interface {
	// Call performs an outbound Raw request.
	Call(req *Request, body []byte) ([]byte, *Response, error)
}

// New builds a new Raw client.
func New(c transport.Channel) Client {
	return rawClient{
		t:       c.Outbound,
		caller:  c.Caller,
		service: c.Service,
		// TODO channel-level TTLs
	}
}

type rawClient struct {
	t transport.Outbound

	caller, service string
}

func (c rawClient) Call(req *Request, body []byte) ([]byte, *Response, error) {
	treq := transport.Request{
		Caller:    c.caller,
		Service:   c.service,
		Encoding:  Encoding,
		Procedure: req.Procedure,
		Headers:   req.Headers,
		Body:      bytes.NewReader(body),
		TTL:       req.TTL, // TODO use default from channel
	}

	tres, err := c.t.Call(req.Context, &treq)
	if err != nil {
		return nil, nil, err
	}
	defer tres.Body.Close()

	resBody, err := ioutil.ReadAll(tres.Body)
	if err != nil {
		return nil, nil, err
	}

	return resBody, &Response{Headers: tres.Headers}, nil
}
