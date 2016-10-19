package main

import (
	"log"
	"sync"
)

// TaskList holds the server's list of tasks that need to be ran or have ran.
type TaskList struct {
	sync.Mutex
	Tasks []Task
}

// NewTaskList creates a new TaskList
func NewTaskList() *TaskList {
	return &TaskList{}
}

// AddTask adds a new task to the list of tasks to be worked on
func (tl *TaskList) AddTask(t Task) {
	tl.Lock()
	tl.Tasks = append(tl.Tasks, t)
	tl.Unlock()
	log.Println("There are now", len(tl.Tasks), "tasks in the Sentinel.")
}

// RemoveTask removes a task from the tasklist
