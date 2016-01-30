package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

func Wait(clients []string, timeout int) {

	begin := time.Now()

	var wg sync.WaitGroup
	cancel := make(chan struct{})

	for _, client := range clients {
		wg.Add(1)
		go func(host string) {
			WaitForHttpRequest(host, cancel)
			wg.Done()
		}(fmt.Sprintf("%s:8080", client))
	}

	timer := time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		close(cancel)
	})

	wg.Wait()

	// There's a race here that might result in assuming that a timeout happend
	// although none happend. It appears when the timer fires after the connection
	// succeeded, but before the check via Stop() below.
	// That shouldn't hap\npen very often and the service was pretty short of timing out

	// anyway, so I guess that's ok for now.
	if !timer.Stop() {
		log.Printf("Error: One or more services timed out after %d second(s)", timeout)
		os.Exit(1)
	}
	fmt.Printf("\nAll services are up after %v!\n", time.Now().Sub(begin))
}

func WaitForHttpRequest(host string, cancel <-chan struct{}) {
	url := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   "/",
	}

	err := errors.New("init")
	for err != nil {
		req, reqErr := http.NewRequest("GET", url.String(), nil)
		if reqErr != nil {
			log.Printf("Warning: Failed to create request for URL '%s' -  skipping service '%s'",
				url.String(), host)
			return
		}

		tr := &http.Transport{}
		client := &http.Client{Transport: tr}
		c := make(chan error, 1)
		go func() {
			_, err := client.Do(req)
			c <- err
		}()

		select {
		case <-cancel:
			tr.CancelRequest(req)
			log.Printf("HTTP: Service %v timed out. Last error: %v", host, err)
			<-c
			return
		case err = <-c:
		}
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Printf("HTTP: Service %v is up\n", host)
}
