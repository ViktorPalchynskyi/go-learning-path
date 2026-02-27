package domain

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrTaskCompleted = errors.New("task completed")
	ErrInvalidTitle = errors.New("invalid title")
	ErrSessionNotFound = errors.New("session not found")
)
