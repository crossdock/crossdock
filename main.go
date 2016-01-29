package main

import (
	"os"
	"strings"
	"time"
)

func main() {

	clients := strings.Split(os.Getenv("XLANG_CLIENTS"), ",")
	servers := strings.Split(os.Getenv("XLANG_SERVERS"), ",")
	behaviors := strings.Split(os.Getenv("XLANG_BEHAVIORS"), ",")

	matrix := Matrix{
		Clients:   clients,
		Servers:   servers,
		Behaviors: behaviors,
	}

	time.Sleep(1 * time.Second)

	results := BeginMatrixTest(matrix)

	OutputResults(results)
}
