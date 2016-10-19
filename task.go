package main

import (
	"log"
	"sync"
	"time"

	"github.com/integrii/sentinel/schedule"
)

// Task represents a task for the sentinel server.
type Task struct {
	sync.Mutex
	Job       Job               // the job to run
	Schedule  schedule.Schedule // the times this task should run
	LastRanAt int64             // the last time this task ran in unix time
	// TODO - capture output of all runs
}

// Run runs the task's job
func (t *Task) Run() error {
	log.Println("Job running:", t.Job)
	t.justRan()
	err := t.Job.Run()
	if err != nil {
		log.Println("Job returned an error:", err)
	}
	return err
}

// justRan sets the LastRanAt time safely
func (t *Task) justRan() {
	t.Lock()
	t.LastRanAt = time.Now().Unix()
	t.Unlock()
}

// NewTask returns a blank new task struct
func NewTask() Task {
	return Task{
		Job:      NewJob(),
		Schedule: schedule.Schedule{},
	}
}
