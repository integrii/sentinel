// Package jobRequest is used to send job requests to a sentinel server
package jobRequest

import (
	"github.com/integrii/sentinel/jobTypes"
	"github.com/integrii/sentinel/schedule"
)

// JobRequest is a set of parameters passed when a task is requested.
type JobRequest struct {
	JobType    jobTypes.Job
	ServerURL  string            // the server url to hit
	Parameters map[string]string // the parameters for this job
	Schedule   schedule.Schedule // the times this job should run
}

// New creates a new job request with the specified parameters.
func New(serverURL string, params map[string]string, s schedule.Schedule) JobRequest {
	jr := NewEmpty()
	jr.ServerURL = serverURL
	jr.Parameters = params
	jr.Schedule = s
	return jr
}

// NewEmpty creates a new job request with no parameters.
func NewEmpty() JobRequest {
	return JobRequest{
		Parameters: make(map[string]string),
	}
}
