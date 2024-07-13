package app

import (
	"backup-cronjob/schedulers"
	"backup-cronjob/tasks"
	"log"
	"time"
)

func loop() {
	count := 0
	for {
		time.Sleep(time.Second)
		count++
		if count > 5 {
			break
		}
	}
}

func Run() {
	log.Println("Starting app ...")
	scheduler := schedulers.Create("TASK 1", tasks.TaskFunction, 2*time.Second)
	scheduler2 := schedulers.Create("TASK 2", tasks.TaskFunction2, 3*time.Second)
	schedulerList := []*schedulers.Scheduler{scheduler, scheduler2}
	schedulers.RunAll(schedulerList)
	loop()
	log.Println("Stopping app ...")
}
