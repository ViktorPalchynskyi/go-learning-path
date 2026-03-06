package service

import (
	"context"
	"fmt"

	"github.com/username/lesson_13/internal/domain"
)



type UserRepo interface {
	Create(ctx context.Context, name, email string) (*domain.User, error)
	FindById(ctx context.Context, id int64) (*domain.User, error)
	FindAll(ctx context.Context) ([]*domain.User, error)
	Delete(ctx context.Context, id int64) error
}

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService{
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, name, email string) (*domain.User, error){
	user, err := s.repo.Create(ctx, name, email)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return user, nil
}

func (s *UserService) FindById(ctx context.Context, id int64) (*domain.User, error){
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return user, nil
}

func (s *UserService) FindAll(ctx context.Context) ([]*domain.User, error){
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return users, nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error{
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}