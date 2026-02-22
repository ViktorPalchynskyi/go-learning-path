package models

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidTitle = errors.New("title not found")
)

type Task struct {
	ID    string
	Title string
}

func NewTask(id, title string) *Task {
	return &Task{ID: id, Title: title}
}