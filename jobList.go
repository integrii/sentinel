package main

import "sync"

// JobList holds the server's list of jobs that need to be ran or have ran.
type JobList struct {
	sync.Mutex
	Jobs []Job
}

// NewJobList creates a new JobList
func NewJobList() JobList {
	return JobList{}
}

// AddJob adds a new job to the list of jobs to be worked on
func (jl *JobList) AddJob(j Job) {
	jl.Lock()
	jl.Jobs = append(jl.Jobs, j)
	jl.Unlock()
}

// RemoveJob removes a job from the joblist TODO
