package repository

import (
	"errors"
	"fmt"

	"github.com/username/lesson_10/internal/domain"
)

var ErrNotFound = errors.New("task not found")

type InMemoryRepo struct {
	tasks map[string]*domain.Task
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{tasks: make(map[string]*domain.Task)}
}

func (r *InMemoryRepo) Save(task *domain.Task) error{
	if task.ID == "" || task.Title == "" {
		return fmt.Errorf("Task is not valid: %v", task)
	}

	r.tasks[task.ID] = task

	return nil
}

func (r *InMemoryRepo) FindById(id string) (*domain.Task, error){
	task := r.tasks[id]
	
	if task == nil {
		return nil, fmt.Errorf("FindById %s: %w", id, ErrNotFound)
	}

	return task, nil
}