package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
)

var listenPort int

func init() {
	flag.IntVar(&listenPort, "port", 80, "Sets the port that the web service listens on.")
}

func main() {
	// the root handler takes new sentinel tasks
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":80", nil)
}

// indexHandler serves the requests for the root url.
// Handles new requests for sentinel tasks.
func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	var body []byte
	var err error

	// Read in the body and return a 400 if unable to
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error reading body from ", req.Host)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the JSON and return a 400 if unable to
	incomingRequest := NewJobRequest()
	err = json.Unmarshal(body, &incomingRequest)
	if err != nil {
		log.Println("Unable to parse JSON from ", req.Host, "had error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// do something with the

}

// JobRequest is a set of parameters passed when a task is requested.
type JobRequest struct {
	JobType    string
	ServerURL  string            // the server url to hit
	Parameters map[string]string // the parameters for this job
	Schedule   Schedule          // the times this job should run
}

// NewJobRequest creatse a new job request with no parameters.
func NewJobRequest() JobRequest {
	return JobRequest{
		Parameters: make(map[string]string),
		Schedule:   Schedule{},
	}
}
