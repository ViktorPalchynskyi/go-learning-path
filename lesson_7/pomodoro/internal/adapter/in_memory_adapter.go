package adapter

import (
	"github.com/username/pomodoro/internal/domain"
	"github.com/username/pomodoro/internal/ports"
)

var _ ports.TaskRepository = (*InMemoryTaskRepo)(nil)

type InMemoryTaskRepo struct {
	tasks map[string]*domain.Task
}

func NewInMemoryTaskRepo() *InMemoryTaskRepo {
	return &InMemoryTaskRepo{tasks: make(map[string]*domain.Task)}
}

func (r *InMemoryTaskRepo) Save(task *domain.Task) error {
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepo) FindById(id string) (*domain.Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return nil, domain.ErrTaskNotFound
	}
	return task, nil
}

func (r *InMemoryTaskRepo) FindAll() ([]*domain.Task, error) {
	result := make([]*domain.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (r *InMemoryTaskRepo) Delete(id string) error {
	if _, ok := r.tasks[id]; !ok {
		return domain.ErrTaskNotFound
	}
	delete(r.tasks, id)
	return nil
}

