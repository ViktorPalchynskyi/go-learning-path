package main

import (
	"fmt"
	"study-go/models"
)

func main() {
	fmt.Println("Lesson 5")
	task1 := models.NewTask("1", "task 1")
	task2 := models.NewTask("2", "task 2")
	task3 := models.NewTask("3", "task 3")

	tasks := []*models.Task{
		task1,
		task2,
		task3,
		models.NewTask("4", "task 4"),
		models.NewTask("5", "task 5"),
		&models.Task{ID: "6", Title: "task 6", Completed: true},
		&models.Task{ID: "7", Title: "task 7", Completed: true},
	}

	for _, t := range filterByCompled(tasks, false) {
		fmt.Printf("%+v\n", t)
	}

	store := make(map[string]*models.Task)
	store["task1"] = task1
	store["task2"] = task2
	store["task3"] = task3
}

func filterByCompled(tasks []*models.Task, completed bool) []*models.Task{
	result := make([]*models.Task, 0, len(tasks))

	for _, t := range tasks {
		if t.Completed == completed {
			result = append(result, t)
		}
	}

	return result
}