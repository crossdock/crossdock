package main

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMatrixFromEnviron(t *testing.T) {

	clients := []string{"yarpc-go", "yarpc-node", "yarpc-browser"}
	servers := []string{"yarpc-go", "yarpc-node"}
	transports := []string{"http", "tchannel"}

	os.Setenv("XLANG_CLIENTS", strings.Join(clients, ","))
	os.Setenv("XLANG_DIMENSION_SERVER", strings.Join(servers, ","))
	os.Setenv("XLANG_DIMENSION_TRANSPORT", strings.Join(transports, ","))

	matrix := ReadMatrixFromEnviron()

	assert.Equal(t, matrix.Clients, clients)
	assert.Equal(t, matrix.Dimensions, []Dimension{
		{Name: "server", Values: servers},
		{Name: "transport", Values: transports},
	})
}
