package schedulers

import (
	"backup-cronjob/tasks"
	"log"
	"time"

	"github.com/algorythma/go-scheduler"
	"github.com/algorythma/go-scheduler/storage"
)

type Scheduler struct {
	TaskStore storage.NoOpStorage
	Scheduler scheduler.Scheduler
	Task      tasks.Task
}

func New(task tasks.Task) *Scheduler {
	return &Scheduler{
		TaskStore: storage.NoOpStorage{},
		Task:      task,
	}
}

func (s *Scheduler) Run() {
	log.Println("Executing task: " + s.Task.Name)
	s.Scheduler.Start()
}

func (s *Scheduler) Stop() {
	log.Println("Stopping task: " + s.Task.Name)
	defer s.Scheduler.Stop()
}

func (s *Scheduler) Configure(interval time.Duration) {
	newScheduler := scheduler.New(s.TaskStore)
	newScheduler.RunEvery(interval, s.Task.Execute)
	s.Scheduler = newScheduler
}
