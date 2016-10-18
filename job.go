package main

import (
	"bytes"
	"log"
	"net/http"
	"net/url"
)

// Job represents some command(s) that need run in a task.
type Job struct {
	JobType    JobType
	URL        string
	Parameters map[string]string
}

// Run runs the job. Returns an error code if it fails.
func (j *Job) Run() error {
	var err error

	// do different things depending on the job type
	switch j.JobType {
	case JOB_HTTPGET:
		// TODO
	case JOB_HTTPPOSTJSON:
		// TODO
		payload := new(bytes.Buffer)
		_, err = http.Post(j.URL, "application/json; charset=utf-8", payload)
	case JOB_HTTPPOST:
		// add all prams as values to the post
		var postValues url.Values
		for name, value := range j.Parameters {
			postValues.Add(name, value)
		}
		_, err := http.PostForm(j.URL, postValues)
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
