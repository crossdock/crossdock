package main

import (
	"os"
	"strings"
)

const clientsKey = "CROSSDOCK_CLIENTS"
const dimensionKeyPrefix = "CROSSDOCK_AXIS_"

// ReadMatrixFromEnviron creates a Matrix by looking for CROSSDOCK_ environment vars
func ReadMatrixFromEnviron() Matrix {
	clients := strings.Split(os.Getenv(clientsKey), ",")
	var dimensions []Dimension

	for _, e := range os.Environ() {
		if !strings.HasPrefix(e, dimensionKeyPrefix) {
			continue
		}
		d := strings.TrimPrefix(e, dimensionKeyPrefix)

		pair := strings.SplitN(d, "=", 2)
		key := strings.ToLower(pair[0])
		value := strings.Split(pair[1], ",")

		dimension := Dimension{
			Name:   key,
			Values: value,
		}
		dimensions = append(dimensions, dimension)
	}

	matrix := Matrix{
		Clients:    clients,
		Dimensions: dimensions,
	}

	return matrix
}
