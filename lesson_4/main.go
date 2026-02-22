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

	tasks := []*models.Task{
		models.NewTask("1", "task 1"),
		models.NewTask("2", "task 2"),
		models.NewTask("3", "task 3"),
		models.NewTask("4", "task 4"),
		models.NewTask("5", "task 5"),
	}

	completeAll(tasks)

	fmt.Println(tasks[0].Completed)
}

func completeAll(tasks []*models.Task)  {
	for _, t := range tasks {
		t.Complete()
	}
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