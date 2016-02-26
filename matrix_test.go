package main

import (
	"testing"

	_ "github.com/orchestrate-io/dvr"
	"github.com/stretchr/testify/assert"
)

func TestExecuteTestCaseWithNon200ShouldFail(t *testing.T) {
	testCase := TestCase{
		Client: "localhost",
		Arguments: Arguments{
			"behavior": "echo",
			"server":   "localhost",
		},
	}
	result := ExecuteCase(testCase)
	assert.Equal(t, Failed, result.Status)
	assert.Equal(t, "Internal Server Error\n", result.Response)

}
