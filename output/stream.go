package output

import (
	"errors"
	"fmt"

	"github.com/yarpc/crossdock/execute"
)

// Stream results to the console, error at end if any fail
func Stream(tests <-chan execute.TestResponse) error {
	failed := false
	for test := range tests {
		for _, result := range test.Results {
			var statStr string
			switch result.Status {
			case execute.Success:
				statStr = "PASSED"
			case execute.Skipped:
				statStr = "SKIPPED"
			default:
				statStr = "FAILED"
				failed = true
			}
			fmt.Printf("%v - %v - %v\n", statStr, test.TestCase, result.Output)
		}
	}
	if failed == true {
		return errors.New("one or more tests failed")
	}
	return nil
}
