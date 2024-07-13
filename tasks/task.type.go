package tasks

import (
	"time"
)

type Task struct {
	Name     string
	Execute  func()
	Interval time.Duration
}

func New(name string, task func(), time time.Duration) *Task {
	return &Task{
		Name:     name,
		Execute:  task,
		Interval: time,
	}
}

func (t *Task) SetName(name string) {
	t.Name = name
}

func (t *Task) SetExecute(function func()) {
	t.Execute = function
}

func (t *Task) SetInterval(interval time.Duration) {
	t.Interval = interval
}
