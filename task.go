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

// UnixTime returns the next unixtime that this task falls under, including this second
// returns a 0 if this task is dead or in the past only
func (t *Task) UnixTime() int64 {

	// the current unix time in seconds
	currentTime := time.Now().Unix()

	// TODO - support reocurring schedules here
	nextSecond := t.Schedule.Time.Unix()
	if currentTime > nextSecond {
		return 0
	}

	return nextSecond
}
