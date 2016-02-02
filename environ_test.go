package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMatrixFromEnviron(t *testing.T) {
	os.Setenv("XLANG_CLIENTS", "yarpc-go,yarpc-node,yarpc-browser")
	os.Setenv("XLANG_DIMENSION_SERVER", "yarpc-go,yarpc-node")
	os.Setenv("XLANG_DIMENSION_TRANSPORT", "http,tchannel")

	matrix := ReadMatrixFromEnviron()

	assert.Equal(t, matrix.Clients, []string{"yarpc-go", "yarpc-node", "yarpc-browser"})
	assert.Equal(t, matrix.Dimensions, []Dimension{
		{Name: "server", Values: []string{"yarpc-go", "yarpc-node"}},
		{Name: "transport", Values: []string{"http", "tchannel"}},
	})
}
