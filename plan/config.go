package plan

import (
	"os"
	"strconv"
	"strings"
	"time"
)

const callDeadlineDefault = 5

// ReadConfigFromEnviron creates a Config by looking for CROSSDOCK_ environment vars
func ReadConfigFromEnviron() *Config {
	const (
		callDeadlineKey = "CROSSDOCK_CALL_DEADLINE"
		clientsKey      = "CROSSDOCK_CLIENTS"
		axisKeyPrefix   = "CROSSDOCK_AXIS_"
	)
	callDeadline, _ := strconv.Atoi(os.Getenv(callDeadlineKey))
	if callDeadline == 0 {
		callDeadline = callDeadlineDefault
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
		CallDeadline: time.Duration(callDeadline),
		Clients:      clients,
		Axes:         axes,
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
