package models

import "fmt"

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

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}