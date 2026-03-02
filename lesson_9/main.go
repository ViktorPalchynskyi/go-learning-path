package main

import (
	"fmt"
	"study-go/models"
	"time"
)

func main()  {
	fmt.Println("Lesson 9")
	cn := make(chan *models.Task, 3)
	ids := []string{"1", "2", "3"}

	for _, id := range(ids) {
		go fetchTask(id, cn)
	}

	for range(ids) {
		task := <-cn
		fmt.Printf("Loaded: %s - %s\n", task.ID, task.Title)
	}
}

func fetchTask(id string, ch chan <- *models.Task) {
	time.Sleep(100 * time.Microsecond)
	ch <- &models.Task{
		ID: id,
		Title: "Task " + id,
	}
}

