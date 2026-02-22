package main

import (
	"fmt"
	"study-go/models"
)

func main() {
	fmt.Println("Lesson 4")

	task1 := models.NewTask("1", "task 1")

	resetTask(*task1)
	fmt.Println(task1.Completed)
	resetTaskPtr(task1)
	fmt.Println(task1.Completed)
}

func resetTask(t models.Task)  {
	t.Completed = !t.Completed
}

func resetTaskPtr(t *models.Task)  {
	t.Completed = !t.Completed
}

func getTaskTitle(t *models.Task) string {
	if t.Title == "" {
		return "untitled"
	}

	return t.Title
}