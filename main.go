package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("\nCrossdock starting...\n\n")
	matrix := ReadMatrixFromEnviron()

	fmt.Printf("Waiting on CROSSDOCK_CLIENTS=%v\n\n", matrix.Clients)
	Wait(matrix.Clients, time.Duration(30)*time.Second)

	fmt.Printf("\nExecuting Matrix...\n\n")
	results := Execute(matrix)

	Output(results)
}
