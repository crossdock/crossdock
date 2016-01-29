package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Matrix struct {
	Clients   []string
	Servers   []string
	Behaviors []string
}

type Results []Result

func BeginMatrixTest(matrix Matrix) Results {

	len := len(matrix.Clients) * len(matrix.Servers) * len(matrix.Behaviors)
	results := make(Results, 0, len)

	for _, client := range matrix.Clients {
		for _, server := range matrix.Servers {
			for _, behavior := range matrix.Behaviors {

				testCase := TestCase{
					Client:   client,
					Server:   server,
					Behavior: behavior,
				}

				result := ExecuteTestCase(testCase)
				results = append(results, result)
			}
		}
	}

	return results
}

type TestCase struct {
	Client   string
	Server   string
	Behavior string
}

type Result struct {
	TestCase TestCase
	Status   Status
	Response string
}

type Status int

const (
	Success Status = 1 + iota
	Failed
	Skipped
)

func ExecuteTestCase(testCase TestCase) Result {

	callUrl, err := url.Parse(fmt.Sprintf("http://%v:8080/", testCase.Client))
	if err != nil {
		log.Fatal(err)
	}

	params := url.Values{}
	params.Add("server", testCase.Server)
	params.Add("behavior", testCase.Behavior)
	callUrl.RawQuery = params.Encode()

	resp, err := http.Get(callUrl.String())
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	status := Success
	if resp.StatusCode != 200 {
		status = Failed
	}

	return Result{
		TestCase: testCase,
		Status:   status,
		Response: string(body),
	}
}
