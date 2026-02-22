package main

import (
	"fmt"
	"study-go/models"
	"time"
)

func main() {
	fmt.Println("Lesson 5")
	task1 := models.NewTask("1", "task 1")
	task2 := models.NewTask("2", "task 2")
	task3 := models.NewTask("3", "task 3")
	task4 := models.NewTask("4", "task 4")
	task5 := models.NewTask("5", "task 5")
	task6 := &models.Task{ID: "6", Title: "task 6", Completed: true}
	task7 := &models.Task{ID: "7", Title: "task 7", Completed: true}

	tasks := []*models.Task{
		task1,
		task2,
		task3,
		task4,
		task5,
		task6,
		task7,
	}

	for _, t := range filterByCompled(tasks, false) {
		fmt.Printf("%+v\n", t)
	}

	store := make(map[string]*models.Task)
	store[task1.ID] = task1
	store[task2.ID] = task2
	store[task3.ID] = task3

	if t, ok := store[task1.ID]; !ok {
		fmt.Printf("error %v", t)
	}

	delete(store, task1.ID)

	for id, t := range store {
		fmt.Printf("%s: %s\n", id, t.Title)
	}

	now := time.Now()
	sessionStore := []*models.Session{
		models.NewWorkSession("1", "1", 25*time.Minute, now.Add(-7*24*time.Hour)),
		models.NewWorkSession("2", "1", 25*time.Minute, now.Add(-6*24*time.Hour)),
		models.NewWorkSession("3", "2", 30*time.Minute, now.Add(-5*24*time.Hour)),
		models.NewWorkSession("4", "2", 15*time.Minute, now.Add(-4*24*time.Hour)),
		models.NewWorkSession("5", "3", 45*time.Minute, now.Add(-3*24*time.Hour)),
		models.NewWorkSession("6", "1", 20*time.Minute, now.Add(-2*24*time.Hour)),
		models.NewWorkSession("7", "3", 25*time.Minute, now.Add(-1*24*time.Hour)),
	}

	for day, sessions := range groupByDay(sessionStore) {
		fmt.Printf("%s: %d sessions\n", day, len(sessions))
		for _, s := range sessions {
			fmt.Printf("  task=%s duration=%s\n", s.TaskID, s.Duration)
		}
	}
}

func groupByDay(sessions []*models.Session) map[string][]*models.Session{
	groups := make(map[string][]*models.Session)

	for _, s := range sessions {
		day := s.StartTime.Format("2006-01-02")
		groups[day] = append(groups[day], s)
	}

	return groups
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