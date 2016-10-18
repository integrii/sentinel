package main

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
