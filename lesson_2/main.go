package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Lesson 2")

	task1 := NewTask("1", "Study Go lesson 1")
	task2 := NewTask("2", "Plan traning")

	fmt.Printf("is task 1 over %v\n", task1.IsCompleted())
	task2.Complete()
	fmt.Printf("is task 2 over %v\n", task2.IsCompleted())	
}

type Task struct {
	ID string
	Title string
	Description string
	Completed bool
	CreatedAt time.Time
}

func NewTask(id, title string) *Task{
	return &Task{
		ID: id,
		Title: title,
		CreatedAt: time.Now(),
	}
}

func (t *Task) Complete()  {
	t.Completed = true
}

func (t *Task) IsCompleted() bool {
	return t.Completed
}