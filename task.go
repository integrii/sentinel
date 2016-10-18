package main

// Task represents a task for the sentinel server.
type Task struct {
	Job      Job      // the job to run
	Schedule Schedule // the times this task should run
	LastRan  int64    // the last time this task ran in unix time
	// TODO - capture output of all runs
}
