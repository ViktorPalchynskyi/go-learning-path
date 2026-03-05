package domain

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNotFound     = errors.New("task not found")
	ErrInvalidTitle = errors.New("invalid title")
)

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func NewTask(id, title string) *Task {
	return &Task{ID: id, Title: title, Completed: false}
}

func (t *Task) Complete() {
	t.Completed = true
}

func ValidateTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return fmt.Errorf("%w: title is required", ErrInvalidTitle)
	}
	if len(title) > 200 {
		return fmt.Errorf("%w: title must be less than 200 characters", ErrInvalidTitle)
	}
	return nil
}
