package main

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecuteToOut(cmd *exec.Cmd) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	if len(output) > 0 {
		fmt.Printf("\n%s\n", string(output))
	}
}
