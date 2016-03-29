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

package plan

import (
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultCallTimeout = 5

// ReadConfigFromEnviron creates a Config by looking for CROSSDOCK_ environment vars
func ReadConfigFromEnviron() *Config {
	const (
		callTimeoutKey = "CROSSDOCK_CALL_TIMEOUT"
		clientsKey     = "CROSSDOCK_CLIENTS"
		axisKeyPrefix  = "CROSSDOCK_AXIS_"
	)
	callTimeout, _ := strconv.Atoi(os.Getenv(callTimeoutKey))
	if callTimeout == 0 {
		callTimeout = defaultCallTimeout
	}
	clients := trimCollection(strings.Split(os.Getenv(clientsKey), ","))
	var axes []Axis
	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, axisKeyPrefix) {
			continue
		}
		d := strings.TrimPrefix(e, axisKeyPrefix)

		pair := strings.SplitN(d, "=", 2)
		key := strings.ToLower(pair[0])
		values := strings.Split(pair[1], ",")
		values = trimCollection(values)

		axis := Axis{
			Name:   key,
			Values: values,
		}
		axes = append(axes, axis)
	}
	config := &Config{
		CallTimeout: time.Duration(callTimeout),
		Clients:     clients,
		Axes:        axes,
	}
	return config
}

func trimCollection(in []string) []string {
	ret := make([]string, len(in))
	for i, v := range in {
		ret[i] = strings.Trim(v, " ")
	}
	return ret
}
