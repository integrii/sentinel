package main

import (
	"bytes"
	"net/http"
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
		payload := new(bytes.Buffer)
		_, err = http.Post(j.URL, "application/json; charset=utf-8", payload)
	}

	return err
}
