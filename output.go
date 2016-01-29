package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func OutputResults(results Results) {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Status", "Client", "Server", "Behavior", "Response"})
	table.SetBorder(false)

	for _, result := range results {

		row := make([]string, 0)

		switch result.Status {
		case Success:
			row = append(row, "PASSED")
		case Failed:
			row = append(row, "FAILED")
		case Skipped:
			row = append(row, "SKIPPED")
		}

		row = append(row, result.TestCase.Client)
		row = append(row, result.TestCase.Server)
		row = append(row, result.TestCase.Behavior)
		row = append(row, result.Response)

		table.Append(row)
	}

	fmt.Println()
	table.Render() // Send output
	fmt.Println()
}
