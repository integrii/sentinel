package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/integrii/sentinel/jobRequest"
)

var listenPort int
var taskList *TaskList // holds all jobs that need running
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
	log.Println("Sentinel listening on port", strconv.Itoa(listenPort))
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe("127.0.0.1:"+strconv.Itoa(listenPort), nil)
	if err != nil {
		log.Println("Error with sentintel web service:", err)
	}
}

// indexHandler serves the requests for the root url.
// Handles new requests for sentinel tasks.
func indexHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("Got a sentinel API request.")
	w.Header().Set("Content-Type", "text/plain")

	var err error

	// Read in the body as json and return a 400 if unable to
	incomingRequest := jobRequest.NewEmpty()
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&incomingRequest)
	if err != nil {
		log.Println("Error decoding job request input:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// convert incoming job request to new job
	log.Println("Adding a new sentinel task.")
	newTask := NewTask()
	newTask.Job.JobType = incomingRequest.JobType
	newTask.Job.Parameters = incomingRequest.Parameters
	newTask.Job.URL = incomingRequest.ServerURL
	newTask.Schedule = incomingRequest.Schedule
	log.Println("New sentinel task being added:", newTask)

	// add the new job to the jobList
	taskList.AddTask(newTask)

	// return a 200 for successful execution
	w.WriteHeader(http.StatusOK)
}
