package schedulers

import (
	"log"
	"time"

	"github.com/algorythma/go-scheduler"
	"github.com/algorythma/go-scheduler/storage"
)

type Scheduler struct {
	Name         string
	TaskStore    storage.NoOpStorage
	Scheduler    scheduler.Scheduler
	TaskFunction func()
}

func New(name string) *Scheduler {
	return &Scheduler{
		Name:      name,
		TaskStore: storage.NoOpStorage{},
	}
}

func (s *Scheduler) Run() {
	log.Println("Executing task: " + s.Name)
	s.Scheduler.Start()
}

func (s *Scheduler) Stop() {
	log.Println("Stopping task: " + s.Name)
	defer s.Scheduler.Stop()
}

func (s *Scheduler) SetConfig(time time.Duration) {
	newScheduler := scheduler.New(s.TaskStore)
	newScheduler.RunEvery(time, s.TaskFunction)
	s.Scheduler = newScheduler
}
