package repository

import "github.com/username/lesson_10/internal/domain"

type InMemoryRepo struct {
	tasks map[string]*domain.Task
}

func NewInMemoryRepo() *InMemoryRepo {
	return &InMemoryRepo{tasks: make(map[string]*domain.Task)}
}
