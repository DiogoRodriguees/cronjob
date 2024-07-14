package app

import (
	"backup-cronjob/schedulers"
	"backup-cronjob/tasks"
	"log"
	"time"
)

func loop() {
	for {
		time.Sleep(time.Second)
	}
}

func Run() {
	log.Println("Starting app ...")
	taskList := tasks.InitilizeTasks()
	schedulerList := schedulers.CreateMany(taskList)
	schedulers.RunMany(schedulerList)
	loop()
	log.Println("Stopping app ...")
}
