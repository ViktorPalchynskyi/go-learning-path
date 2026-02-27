package main

import (
	"fmt"

	"github.com/username/pomodoro/internal/adapter"
	"github.com/username/pomodoro/internal/service"
)

func main()  {
	fmt.Println("Pomodoro server starting...")

	repo := adapter.NewInMemoryTaskRepo()
	srv := service.NewTaskService(repo)
	task, err := srv.CreateTask("Test")

	if err != nil {
        fmt.Errorf("error %w", err)
    }
    fmt.Printf("Created: %s - %s\n", task.ID, task.Title)
}