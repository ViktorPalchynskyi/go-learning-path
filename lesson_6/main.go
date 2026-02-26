package main

import (
	"errors"
	"fmt"
	"strings"
	"study-go/models"
)

func main() {
	fmt.Println("Lesson 6")

	repo := models.NewInMemoryTaskRepo()
	taskService := models.NewTaskService(repo)
	task, err := taskService.GetTask("1")
	if err != nil {
		fmt.Printf("error getting task: %v\n", err)
	}

	fmt.Printf("task: %+v\n", task)

	task, err = taskService.CreateTask("")
	err = models.ValidateTask(task)
	var ve *models.ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("validation error: %s: %s\n", ve.Field, ve.Message)
	}
}

var (
	ErrTaskNotFound = errors.New("task not found")
	ErrTaskCompleted = errors.New("task completed")
	ErrInvalidTitle = errors.New("invalid title")
	ErrSessionNotFound = errors.New("session not found")
)

func validateTitle(title string) error {
	if strings.TrimSpace(title) == "" {
		return fmt.Errorf("%w: title is required", ErrInvalidTitle)
	}

	if len(title) > 200 {
		return fmt.Errorf("%w: title must be less than 200 characters", ErrInvalidTitle)
	}

	return nil
}