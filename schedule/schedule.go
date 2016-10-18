// Package schedule implements a schedule used by sentinel tasks
package schedule

import "time"

// Schedule is a set of times in which a task should run.
type Schedule struct {
	Time               time.Time
	RepeatSeconds      []int
	RepeatMinutes      []int
	RepeatHours        []int
	RepeatDaysOfWeek   []int
	RepeatDaysOfMonth  []int
	RepeatMonthsOfYear []int
	// TODO - add times to skip in the same way?
}

// NewOneTimeSchedule returns a new schedule that runs only once
// at the specified time
func NewOneTimeSchedule(t time.Time) Schedule {
	s := New()
	s.Time = t
	return s
}

// New returns a new blank schedule struct
func New() Schedule {
	return Schedule{}
}
