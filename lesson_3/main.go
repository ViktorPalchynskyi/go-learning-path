package main

import (
	"fmt"
	"study-go/models"
)

func main() {
	fmt.Println("Lesson 3")
}

type TaskRepository interface {
	Save(task *models.Task) error
	FindById(id string) (*models.Task, error)
	FindAll() ([]*models.Task, error)
	Delete(id string) error
}

type InMemoryTaskRepo struct {
	tasks map[string]*models.Task
}

func NewInMemoryTaskRepo() *InMemoryTaskRepo{
    return &InMemoryTaskRepo{tasks: make(map[string]*models.Task)}
}

// Delete implements [TaskRepository].
func (i *InMemoryTaskRepo) Delete(id string) error {
    if _, ok := i.tasks[id];!ok {
        return models.ErrTaskNotFound
    }

    delete(i.tasks, id)
    
    return nil
}

// FindAll implements [TaskRepository].
func (i *InMemoryTaskRepo) FindAll() ([]*models.Task, error) {
	result := make([]*models.Task, 0, len(i.tasks))

    for _, t := range i.tasks {
        result = append(result, t)
    }

    return result, nil
}

// FindById implements [TaskRepository].
func (i *InMemoryTaskRepo) FindById(id string) (*models.Task, error) {
    task, ok := i.tasks[id]

    if !ok {
        return nil, models.ErrTaskNotFound
    }

	return task, nil
}

// Save implements [TaskRepository].
func (i *InMemoryTaskRepo) Save(task *models.Task) error {
	i.tasks[task.ID] = task
    return nil
}

var _ TaskRepository = (*InMemoryTaskRepo)(nil)
