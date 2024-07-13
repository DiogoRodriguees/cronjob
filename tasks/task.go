package tasks

import "time"

func InitilizeTasks() []Task {
	task1 := New("TASK 1", TaskOne, 2*time.Second)
	task2 := New("TASK 2", TaskTwo, 3*time.Second)

	return []Task{*task1, *task2}
}
