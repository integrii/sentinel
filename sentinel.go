package main

import (
	"log"
	"time"
)

// Sentinel watches for jobs that need ran and delegates their tasks to run
type Sentinel struct {
	LastProcessedSecond int64 // the last second that was processed by the system
	JobsRun             int   // the number of jobs that have been fired
}

// RunForTime runs jobs for the specified number of seconds (or before if missed)
func (s *Sentinel) RunForTime(unixTime int64) {
	for _, task := range taskList.Tasks {
		if task.UnixTime() == unixTime {
			go task.Run()
		}
	}
}

// Run runs the sentinel and spawns go procs for each unix second
// each spawned processes looks for tasks and spawns any tasks
// that are ready to run as goroutines.  Runs forever.
func (s *Sentinel) Run() {
	log.Println("Sentinel running.")
	for {
		var currentSecond = time.Now().Unix()
		// if its time to run again, scan for jobs that need ran
		for currentSecond > s.LastProcessedSecond {
			go s.RunForTime(currentSecond)
			log.Println("Sentinel running for second", currentSecond)
			s.LastProcessedSecond = currentSecond
		}
	}
	log.Println("Sentinel stopped.")
}
