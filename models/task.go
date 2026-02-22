package models

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrInvalidTitle = errors.New("title not found")
)

type Task struct {
	ID    string
	Title string
	Completed bool
}

func NewTask(id, title string) *Task {
	return &Task{ID: id, Title: title, Completed: false}
}

func (t *Task) SafeTitle() string{
	if t == nil {
		return "untitled"
	}

	return t.Title
}