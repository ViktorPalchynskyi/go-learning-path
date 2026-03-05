package service

import (
	"errors"
	"testing"

	"github.com/username/lesson_11/internal/domain"
)

type mockTaskRepo struct {
	tasks map[string]*domain.Task
}

func newMockRepo() *mockTaskRepo {
	return &mockTaskRepo{tasks: make(map[string]*domain.Task)}
}

func (r *mockTaskRepo) Save(task *domain.Task) error {
	r.tasks[task.ID] = task
	return nil
}

func (r *mockTaskRepo) FindById(id string) (*domain.Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return task, nil
}

func (r *mockTaskRepo) FindAll() ([]*domain.Task, error) {
	result := make([]*domain.Task, 0, len(r.tasks))
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (r *mockTaskRepo) Update(task *domain.Task) error {
	if _, ok := r.tasks[task.ID]; !ok {
		return domain.ErrNotFound
	}
	r.tasks[task.ID] = task
	return nil
}

func (r *mockTaskRepo) Delete(id string) error {
	if _, ok := r.tasks[id]; !ok {
		return domain.ErrNotFound
	}
	delete(r.tasks, id)
	return nil
}

func TestTaskService_CreateTask(t *testing.T) {
	svc := NewTaskService(newMockRepo())

	task, err := svc.CreateTask("Learn Go")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if task.Title != "Learn Go" {
		t.Errorf("title = %q, want %q", task.Title, "Learn Go")
	}
	if task.ID == "" {
		t.Error("ID is empty, want non-empty")
	}
}

func TestTaskService_GetTask_NotFound(t *testing.T) {
	svc := NewTaskService(newMockRepo())

	_, err := svc.GetTask("nonexistent")
	if !errors.Is(err, domain.ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}
