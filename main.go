package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/integrii/sentinel/jobRequest"
)

var listenPort int
var taskList TaskList // holds all jobs that need running
var sentinel Sentinel
var sigChan chan struct{}

func init() {
	flag.IntVar(&listenPort, "port", 80, "Sets the port that the web service listens on.")
	taskList = NewTaskList()
	sentinel = Sentinel{}
	sigChan = make(chan struct{})
}

func main() {
	go startWebServer()
	go sentinel.Run()
	<-sigChan
}

// startWebServer starts the web server for incoming api requests
func startWebServer() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+string(listenPort), nil)
}

// indexHandler serves the requests for the root url.
// Handles new requests for sentinel tasks.
func indexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	var err error

	// Read in the body as json and return a 400 if unable to
	decoder := json.NewDecoder(req.Body)
	incomingRequest := jobRequest.NewEmpty()
	err = decoder.Decode(incomingRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert incoming job request to new job
	newTask := NewTask()
	newTask.Job.JobType = JobType(incomingRequest.JobType)
	newTask.Job.Parameters = incomingRequest.Parameters
	newTask.Job.URL = incomingRequest.ServerURL

	// add the new job to the jobList
	taskList.AddTask(newTask)

	// return a 200 for successful execution
	w.WriteHeader(http.StatusOK)
}
