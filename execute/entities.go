package execute

import "github.com/yarpc/crossdock/plan"

// TestResponse contains the reply from a test client, most importantly,
// contains a list of Results for the test cases ran by the test client
type TestResponse struct {
	TestCase plan.TestCase
	Results  []Result
}

// Result is the outcome of an individual test case ran by the test client
type Result struct {
	Status Status
	Output string
}

// Status is an enum that represents test success/failure
type Status int

const (
	// Success indicates a client's TestCase passed
	Success Status = 1 + iota

	// Failed indicates a client's TestCase did not pass
	Failed

	// Skipped indicates a client' TestCase did not run
	Skipped
)
