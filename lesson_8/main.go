package main

import (
	"encoding/json"
	"fmt"
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
}