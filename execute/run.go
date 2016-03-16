package execute

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"

	"github.com/yarpc/crossdock/plan"
)

// Run the test program for a given Plan
func Run(plan *plan.Plan) <-chan TestResponse {
	tests := make(chan TestResponse, 100)
	go func() {
		for _, c := range plan.TestCases {
			tests <- executeTestCase(c)
		}
		close(tests)
	}()
	return tests
}

func executeTestCase(testCase plan.TestCase) TestResponse {
	response, err := makeRequest(testCase)
	if err != nil {
		return TestResponse{
			TestCase: testCase,
			Results: []Result{{
				Status: Failed,
				Output: fmt.Sprintf("err: %v", err),
			}},
		}
	}

	var subResponses []subResponse
	if err := json.Unmarshal(response, &subResponses); err != nil {
		return TestResponse{
			TestCase: testCase,
			Results: []Result{{
				Status: Failed,
				Output: fmt.Sprintf("err: %v", err),
			}},
		}
	}

	return TestResponse{
		TestCase: testCase,
		Results:  toResults(subResponses),
	}
}

func makeRequest(testCase plan.TestCase) ([]byte, error) {
	callURL, err := url.Parse(fmt.Sprintf("http://%v:8080/", testCase.Client))
	if err != nil {
		return []byte(""), err
	}

	args := url.Values{}
	for k, v := range testCase.Arguments {
		args.Add(k, v)
	}
	callURL.RawQuery = args.Encode()

	ctx, _ := context.WithTimeout(
		context.Background(), testCase.Plan.Config.CallDeadline*time.Second)
	resp, err := ctxhttp.Get(ctx, nil, callURL.String())
	if err != nil {
		return []byte(""), err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	if resp.StatusCode != 200 {
		return []byte(""), fmt.Errorf("wanted status code 200, got %v", resp.StatusCode)
	}

	return body, nil
}

type subResponse struct {
	Status string
	Output string
}

func toResults(subResponses []subResponse) []Result {
	var results []Result
	for _, subResponse := range subResponses {
		status := Failed
		switch subResponse.Status {
		case "passed":
			status = Success
		case "skipped":
			status = Skipped
		}
		result := Result{
			Status: status,
			Output: subResponse.Output,
		}
		results = append(results, result)
	}
	return results
}
