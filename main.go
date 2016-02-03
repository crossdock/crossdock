package main

import (
	"fmt"
	"time"
)

func main() {
	matrix := ReadMatrixFromEnviron()

	fmt.Printf("Waiting on XLANG_CLIENTS=%v\n\n", matrix.Clients)
	Wait(matrix.Clients, time.Duration(30)*time.Second)

	fmt.Println("Begining matrix of tests")
	results := Execute(matrix)

	Output(results)
}
