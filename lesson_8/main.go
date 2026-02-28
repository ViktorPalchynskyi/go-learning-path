package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Lesson 8")

	task := &Task{
		ID: "123",
		Title: "Test1",
	}
	data, err := json.Marshal(task)
	if err != nil {
		fmt.Errorf("%w", err)
	}
	fmt.Println(string(data))
}

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description,omitempty"`
	Completed bool `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateTaskRequest struct {
	Title string `json:"title"`
	Description string `json:"description,omitempty"`
}

type TaskResponse struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
	CreatedAt string `json:"created_at"`
}

func toTaskReponse(t *Task) TaskResponse {
	return TaskResponse{
		ID: t.ID,
		Title: t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt.Format(time.RFC3339),
	}
}