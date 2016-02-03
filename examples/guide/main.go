package main

import "net/http"

func main() {
	// xlang makes all calls to http://<test-client>:8080/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// custom arguments, called dimensions, are configured in
		// docker-compose.yml and then passed as query params like so:
		// http://<test-client>:8080/?behavior=dance
		behavior := r.FormValue("behavior")

		// when client is called with no arguments,
		// report back with a 200 when ready to run tests
		if behavior == "" {
			return
		}

		// once the client is ready, xlang will make an HTTP request
		// to / with dimensions you defined in docker-compose.yml,
		// in this case, we've defined XLANG_DIMENSION_BEHAVIOR
		switch behavior {

		// we've recieved a request asking us to test the "dance" behavior,
		// write whatever code we need to verify that behavior, then
		// respond in TAP (testanything.org) format: simply "ok" or "not ok"
		case "dance":
			w.Write([]byte("ok\n"))
			return
		}

		// give a 404 when no behavior defined,
		// xlang will mark every 404 test case as "skipped"
		http.NotFound(w, r)
	})
	http.ListenAndServe(":8080", nil)
}
