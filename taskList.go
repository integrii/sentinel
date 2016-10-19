package main

import (
	"errors"
	"log"
	"sync"
)

// TaskList holds the server's list of tasks that need to be ran or have ran.
type TaskList struct {
	sync.Mutex
	Tasks map[string]Task
}

// NewTaskList creates a new TaskList
func NewTaskList() *TaskList {
	return &TaskList{
		Tasks: make(map[string]Task),
	}
}

// AddTask adds a new task to the list of tasks to be worked on
func (tl *TaskList) AddTask(t Task) {
	tl.Lock()
	tl.Tasks[t.ID] = t
	tl.Unlock()
	log.Println("There are now", len(tl.Tasks), "tasks in the Sentinel.")
}

// RemoveTask removes a task from the tasklist
func (tl *TaskList) RemoveTask(id string) error {
	var err error

	// make sure the key exists in the map before deleting
	_, ok := tl.Tasks[id]
	if !ok {
		return errors.New("That task does not exist.")
	}

	// delete the key from the map
	delete(tl.Tasks, id)

	return err
}
