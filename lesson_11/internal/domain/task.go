package domain

import "errors"

var ErrNotFound = errors.New("task not found")

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTask(id, title string) *Task {
	return &Task{ID: id, Title: title, Completed: false}
}
