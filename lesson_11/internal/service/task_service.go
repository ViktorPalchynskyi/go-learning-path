package service

import (
	"github.com/google/uuid"
	"github.com/username/lesson_11/internal/domain"
)

type TaskRepository interface {
	Save(task *domain.Task) error
	FindById(id string) (*domain.Task, error)
	FindAll() ([]*domain.Task, error)
	Update(task *domain.Task) error
	Delete(id string) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(title string) (*domain.Task, error) {
	task := domain.NewTask(uuid.New().String(), title)
	if err := s.repo.Save(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) GetTask(id string) (*domain.Task, error) {
	return s.repo.FindById(id)
}

func (s *TaskService) ListTasks() ([]*domain.Task, error) {
	return s.repo.FindAll()
}

func (s *TaskService) UpdateTask(id, title string, completed bool) (*domain.Task, error) {
	task, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	task.Title = title
	task.Completed = completed
	if err := s.repo.Update(task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *TaskService) DeleteTask(id string) error {
	return s.repo.Delete(id)
}
