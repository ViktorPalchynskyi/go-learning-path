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
	task1.Touch()
	fmt.Printf("full struct %#v", task1)
}

type BaseEntity struct {
	ID string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *BaseEntity) Touch() {
	b.UpdatedAt = time.Now()
}

type Task struct {
	BaseEntity
	Title string
	Description string
	Completed bool
}

func NewTask(id, title string) *Task{
	return &Task{
		Title: title,
	}
}

func (t *Task) Complete()  {
	t.Completed = true
}

func (t *Task) IsCompleted() bool {
	return t.Completed
}

type SessionType string

const (
	WorkSession SessionType = "work"
	BreakSession SessionType = "break"
)

type Session struct {
	BaseEntity
	TaskID string
	StartTime time.Time
	EndTime time.Time
	Duration time.Duration
	Type SessionType 
}

func NewWorkSession(id, taskId string, duration time.Duration) *Session{
	return &Session{
		TaskID: taskId,
		Duration: duration,
	}
}

func (s *Session) Finish(){
	s.EndTime = time.Now()
}

func (s *Session) IsRunning() bool{
	return s.EndTime.IsZero()
}