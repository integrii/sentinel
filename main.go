package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/integrii/sentinel/jobRequest"
)

var listenPort int
var listenIP string
var taskList *TaskList // holds all jobs that need running
var sentinel Sentinel
var sigChan chan struct{}

func init() {
	flag.IntVar(&listenPort, "port", 80, "Sets the port that the web service listens on.")
	taskList = NewTaskList()
	sentinel = Sentinel{}
	listenIP = "0.0.0.0"
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
	http.HandleFunc("/delete", deleteHandler)
	err := http.ListenAndServe(listenIP+":"+strconv.Itoa(listenPort), nil)
	if err != nil {
		log.Println("Error with sentintel web service:", err)
	}
}

// deleteHandler handles requests for task deletions
func deleteHandler(w http.ResponseWriter, req *http.Request) {
	// Fetch the task ID from the body to delete
	taskID, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Error reading request body for deletion request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = taskList.RemoveTask(string(taskID))
	if err != nil {
		log.Println("Error removing a task", taskID, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

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
	w.Write([]byte(newTask.ID))
	w.WriteHeader(http.StatusOK)
}
