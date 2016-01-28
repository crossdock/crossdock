package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {

	clients := strings.Split(os.Getenv("XLANG_CLIENTS"), ",")
	servers := strings.Split(os.Getenv("XLANG_SERVERS"), ",")
	behaviors := strings.Split(os.Getenv("XLANG_BEHAVIORS"), ",")

	matrix := Matrix{
		Clients:   clients,
		Servers:   servers,
		Behaviors: behaviors,
	}

	beginMatrixTest(matrix)
}

type Matrix struct {
	Clients   []string
	Servers   []string
	Behaviors []string
}

func beginMatrixTest(matrix Matrix) {

	for _, client := range matrix.Clients {
		for _, server := range matrix.Servers {
			for _, behavior := range matrix.Behaviors {

				testCase := TestCase{
					Client:   client,
					Server:   server,
					Behavior: behavior,
				}

				executeTestCase(testCase)
			}
		}
	}
}

type TestCase struct {
	Client   string
	Server   string
	Behavior string
}

func executeTestCase(testCase TestCase) {

	// start with base url
	callUrl, err := url.Parse(fmt.Sprintf("http://%v:8080/", testCase.Client))
	if err != nil {
		log.Fatal(err)
	}

	// build query params
	params := url.Values{}
	params.Add("server", testCase.Server)
	params.Add("behavior", testCase.Behavior)
	callUrl.RawQuery = params.Encode()

	// make request to test client
	resp, err := http.Get(callUrl.String())
	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp.StatusCode, callUrl.String())
}
