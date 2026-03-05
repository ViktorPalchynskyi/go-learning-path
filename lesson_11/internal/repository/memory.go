package repository

import (
	"fmt"

	"github.com/username/lesson_11/internal/domain"
)

type InMemoryRepo struct {
	tasks map[string]*domain.Task
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{tasks: make(map[string]*domain.Task)}
}

func (r *InMemoryRepo) Save(task *domain.Task) error {
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryRepo) FindById(id string) (*domain.Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return nil, fmt.Errorf("FindById %s: %w", id, domain.ErrNotFound)
	}
	return task, nil
}

func (r *InMemoryRepo) FindAll() ([]*domain.Task, error) {
	result := make([]*domain.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (r *InMemoryRepo) Update(task *domain.Task) error {
	if _, ok := r.tasks[task.ID]; !ok {
		return fmt.Errorf("Update %s: %w", task.ID, domain.ErrNotFound)
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryRepo) Delete(id string) error {
	if _, ok := r.tasks[id]; !ok {
		return fmt.Errorf("Delete %s: %w", id, domain.ErrNotFound)
	}
	delete(r.tasks, id)
	return nil
}
