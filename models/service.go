package models

import (
	"fmt"
	"github.com/google/uuid"
)

var _ TaskRepository = (*InMemoryTaskRepo)(nil)

type TaskService struct {
    repo TaskRepository
}

func (s *TaskService) GetTask(id string)(*Task, error) {
	task, err := s.repo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("getting task: %s: %w", id, err)
	}

	return task, nil
}

func (s *TaskService) CreateTask(title string) (*Task, error) {                                                                                                                             
	task := NewTask(uuid.New().String(), title)                                                                                                                                           
	err := s.repo.Save(task)
	if err != nil {
			return nil, fmt.Errorf("creating task: %w", err)
	}
	return task, nil
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}