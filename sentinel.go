package main

import (
	"log"
	"time"
)

// SentinelWorkers is the number of go routines that will
// run jobs from the job task queue
const SentinelWorkers = 50

// Sentinel watches for jobs that need ran and delegates their tasks to run
type Sentinel struct {
	LastProcessedSecond int64 // the last second that was processed by the system
	JobsRun             int   // the number of jobs that have been fired
}

// RunForTime runs jobs for the specified number of seconds (or before if missed)
func (s *Sentinel) RunForTime(unixTime int64, taskQueue chan<- Task) {
	// log.Println("Sentinel checking", len(taskList.Tasks), "tasks.")
	for _, task := range taskList.Tasks {
		// TODO - support reocurring timelines here and not just a single time
		if task.Schedule.Time.Unix() == unixTime {
			taskQueue <- task
		}
	}
}

// Run runs the sentinel and spawns go procs for each unix second
// each spawned processes looks for tasks and spawns any tasks
// that are ready to run as goroutines.  Runs forever.
func (s *Sentinel) Run() {
	// Start the task runner pool
	taskQueue := make(chan Task)
	for i := 0; i <= SentinelWorkers; i++ {
		go s.StartWorker(taskQueue)
	}
	defer close(taskQueue)

	log.Println("Sentinel running.")
	for {
		var currentSecond = time.Now().Unix()
		// if its time to run again, scan for jobs that need ran
		for currentSecond > s.LastProcessedSecond {
			go s.RunForTime(currentSecond, taskQueue)
			log.Println("Sentinel running for second", currentSecond)
			s.LastProcessedSecond = currentSecond
		}
	}
	log.Println("Sentinel stopped.")

}

// StartWorker starts a task worker that closes when the taskQueue channel closes
func (s *Sentinel) StartWorker(taskQueue <-chan Task) {
	for task := range taskQueue {
		task.Run()
		// TODO - check task output and handle here
	}
}
