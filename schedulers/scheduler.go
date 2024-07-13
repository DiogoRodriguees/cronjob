package schedulers

import (
	"time"
)

func Create(name string, taskFunction func(), time time.Duration) *Scheduler {
	scheduler := New(name)
	scheduler.TaskFunction = taskFunction
	scheduler.SetConfig(time)
	return scheduler
}

func RunAll(schedulers []*Scheduler) {
	for _, scheduler := range schedulers {
		scheduler.Run()
	}
}
