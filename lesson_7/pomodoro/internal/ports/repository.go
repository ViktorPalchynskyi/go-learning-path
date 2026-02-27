package ports

import "github.com/username/pomodoro/internal/domain"

type TaskRepository interface {
	Save(task *domain.Task) error
	FindById(id string) (*domain.Task, error)
	FindAll() ([]*domain.Task, error)
	Delete(id string) error
}