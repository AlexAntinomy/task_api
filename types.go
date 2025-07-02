package main

import (
	"time"
)

type TaskStatus string

const (
	StatusPending  TaskStatus = "pending"
	StatusRunning  TaskStatus = "running"
	StatusDone     TaskStatus = "done"
	StatusError    TaskStatus = "error"
	StatusCanceled TaskStatus = "canceled"
)

type Task struct {
	ID        string     `json:"id"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	Duration  float64    `json:"duration"`
	Result    string     `json:"result,omitempty"`
	Err       string     `json:"error,omitempty"`
}
