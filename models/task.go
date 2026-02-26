package models

import (
	"errors"
	"fmt"
)

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


func (t *Task) Complete()  {
	t.Completed = true
}

type ValidationError struct {
	Field string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func ValidateTask(t *Task) error {
	if t.Title == "" {
		return &ValidationError{Field: "title", Message: "title is required"}
	}

	return nil
}