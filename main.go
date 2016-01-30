package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	fmt.Printf("\nXLANG COMMENCING...\n\n")

	clients := strings.Split(os.Getenv("XLANG_CLIENTS"), ",")
	servers := strings.Split(os.Getenv("XLANG_SERVERS"), ",")
	behaviors := strings.Split(os.Getenv("XLANG_BEHAVIORS"), ",")

	// wait on deps to start. use lib directly when ready:
	// @see https://github.com/Barzahlen/waitforservices/issues/4
	fmt.Printf("Waiting on XLANG_CLIENTS=%v\n", clients)
	ExecuteToOut(exec.Command("waitforservices", "-httpport=8080"))

	matrix := Matrix{
		Clients:   clients,
		Servers:   servers,
		Behaviors: behaviors,
	}

	fmt.Println("Beginning Matrix Test:")
	results := BeginMatrixTest(matrix)

	OutputResults(results)
}
