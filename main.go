package main

import "fmt"

func main() {
	matrix := ReadMatrixFromEnviron()

	fmt.Printf("Waiting on XLANG_CLIENTS=%v\n\n", matrix.Clients)
	Wait(matrix.Clients, 30)

	fmt.Println("Begining matrix of tests")
	results := Execute(matrix)

	Output(results)
}
