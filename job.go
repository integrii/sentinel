package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/integrii/sentinel/jobTypes"
)

// Job represents some command(s) that need run in a task.
type Job struct {
	JobType    jobTypes.Job
	URL        string
	Parameters map[string]string
}

// Run runs the job. Returns an error code if it fails.
func (j *Job) Run() error {
	var err error

	// do different things depending on the job type
	switch j.JobType {
	case jobTypes.HTTPGET:
		// TODO
	case jobTypes.HTTPPOST:
		// add all prams as values to the post
		var postValues = url.Values{}
		for name, value := range j.Parameters {
			postValues.Add(name, value)
		}
		resp, err := http.PostForm(j.URL, postValues)
		log.Println("Sending values to URL", j.URL, postValues)
		log.Println("Job run got status code return", resp.StatusCode)
		if err != nil {
			log.Println("Had an error when sending POST", err)
		}
	}

	// flag job as ran

	return err
}

// NewJob creates a new empty job struct
func NewJob() Job {
	return Job{
		Parameters: make(map[string]string),
	}
}
