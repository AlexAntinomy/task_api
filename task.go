package main

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	tasks = map[string]*Task{}
	mu    sync.Mutex
)

func NewTask() *Task {
	t := &Task{
		ID:        uuid.NewString(),
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
	mu.Lock()
	tasks[t.ID] = t
	mu.Unlock()
	go runTask(t)
	return t
}

func runTask(t *Task) {
	mu.Lock()
	t.Status = StatusRunning
	mu.Unlock()
	start := time.Now()
	time.Sleep(3 * time.Minute)
	mu.Lock()
	t.Status = StatusDone
	t.Duration = time.Since(start).Seconds()
	t.Result = "ok"
	mu.Unlock()
}

func GetTask(id string) *Task {
	mu.Lock()
	defer mu.Unlock()
	return tasks[id]
}

func DeleteTask(id string) {
	mu.Lock()
	defer mu.Unlock()
	delete(tasks, id)
}
