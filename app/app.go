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
	tasks := tasks.InitilizeTasks()
	schedulerList := schedulers.CreateMany(tasks)
	schedulers.RunMany(schedulerList)
	loop()
	log.Println("Stopping app ...")
}
