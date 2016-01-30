package main

import (
	"fmt"
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

	fmt.Println("Waiting 1 second for test clients to come online")
	time.Sleep(1 * time.Second)

	fmt.Println("Begining matrix of tests")
	results := BeginMatrixTest(matrix)

	OutputResults(results)
}
