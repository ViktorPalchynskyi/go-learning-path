package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/username/lesson_10/internal/domain"
)

type TaskRepository interface {
	Save(task *domain.Task) error
	FindById(id string) (*domain.Task, error)
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (ts *TaskService) GetTask(ctx context.Context, id string) (*domain.Task, error){
	select{
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	task, err := ts.repo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("GetTask %s: %w", id, err)
	}

	return task, nil
}

func (ts *TaskService) CreateTask(ctx context.Context, title string) (*domain.Task, error){
	select {
	case <- ctx.Done():
		return nil, ctx.Err()
	default:
	}


	task := domain.NewTask(uuid.New().String(), title)
	err := ts.repo.Save(task)

	if err != nil {
		return nil, fmt.Errorf("CreateTask %s: %w", task.ID, err)
	}

	return task, nil
}