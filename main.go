package main

import (
	"fmt"
	"time"
)

func main() {
	matrix := ReadMatrixFromEnviron()

	fmt.Println("Waiting 1 second for test clients to come online")
	time.Sleep(1 * time.Second)

	fmt.Println("Begining matrix of tests")
	results := Execute(matrix)

	Output(results)
}
