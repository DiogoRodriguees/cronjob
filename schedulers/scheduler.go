package schedulers

import (
	"backup-cronjob/tasks"
)

func Create(task tasks.Task) *Scheduler {
	scheduler := New(task)
	scheduler.Configure(task.Interval)
	return scheduler
}

func RunMany(schedulers []*Scheduler) {
	for _, scheduler := range schedulers {
		scheduler.Run()
	}
}

func CreateMany(tasks []tasks.Task) []*Scheduler {
	var schedulerList []*Scheduler
	for _, task := range tasks {
		Create(task)
	}
	return schedulerList
}
