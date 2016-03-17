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
