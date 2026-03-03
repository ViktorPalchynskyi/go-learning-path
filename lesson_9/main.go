package main

import (
	"errors"
	"fmt"
	"study-go/models"
	"sync"
	"time"
)

func main()  {
	fmt.Println("Lesson 9")
	cn := make(chan *models.Task, 3)
	ids := []string{"1", "2", "3"}
	tasks := make([]*models.Task, 10)
	for i := range tasks {
		tasks[i] = models.NewTask(fmt.Sprintf("%d", i), fmt.Sprintf("Task %d", i))
	}

	var wg sync.WaitGroup

	for _, t := range tasks {
		wg.Add(1)
		go func (task *models.Task)  {
			defer wg.Done()
			time.Sleep(100 * time.Microsecond)
			task.Complete()
		}(t)
	}
	wg.Wait()

	fmt.Println("All tasks completed")
	for _, id := range(ids) {
		go fetchTask(id, cn)
	}

	for range(ids) {
		task := <-cn
		fmt.Printf("Loaded: %s - %s\n", task.ID, task.Title)
	}
	task1, err := fetchWithTimeout("1", 200)
	if err != nil {
		fmt.Printf("Error task1: %s\n", err)
	} else {
		fmt.Printf("Loaded: %s - %s\n", task1.ID, task1.Title)
	}

	task2, err := fetchWithTimeout("2", 50)
	if err != nil {
		fmt.Printf("Error task2: %s\n", err)
	} else {
		fmt.Printf("Loaded: %s - %s\n", task2.ID, task2.Title)
	}

}

func fetchTask(id string, ch chan <- *models.Task) {
	time.Sleep(100 * time.Microsecond)
	ch <- &models.Task{
		ID: id,
		Title: "Task " + id,
	}
}

func fetchWithTimeout(id string, timeout time.Duration) (*models.Task, error){
	cn := make(chan *models.Task,1)

	go func ()  {
		time.Sleep(100 * time.Millisecond)
		cn <-&models.Task{
			ID: id,
			Title: "Task " + id,
		}
	}()

	select {
	case task := <-cn:
		return task, nil
	case <-time.After(timeout * time.Millisecond):
		fmt.Println("Timeout")
		return nil, errors.New("Timeout")
	}
}
