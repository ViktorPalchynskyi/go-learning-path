package service

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/username/pomodoro/internal/domain"
	"github.com/username/pomodoro/internal/ports"
)

type TaskService struct {
    repo ports.TaskRepository
}

func (s *TaskService) CreateTask(title string) (*domain.Task, error) {                                                                                                                             
	task := domain.NewTask(uuid.New().String(), title)                                                                                                                                           
	err := s.repo.Save(task)
	if err != nil {
			return nil, fmt.Errorf("creating task: %w", err)
	}
	return task, nil
}

func NewTaskService(repo ports.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}