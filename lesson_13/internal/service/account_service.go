package service

import (
	"context"
	"fmt"

	"github.com/username/lesson_13/internal/domain"
)

type AccountRepo interface {
	Create(ctx context.Context, ownerID int64, balance float64) (*domain.Account, error)
	FindById(ctx context.Context, id int64) (*domain.Account, error)
	Transfer(ctx context.Context, fromID, toID int64, amount float64) error
}

type AccountService struct {
	repo AccountRepo
}

func NewAccountService(repo AccountRepo) *AccountService {                                               
	return &AccountService{repo: repo}
}


func (s *AccountService)Create(ctx context.Context, ownerID int64, balance float64) (*domain.Account, error){
	account, err := s.repo.Create(ctx, ownerID, balance)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return account, nil
}

func (s *AccountService) FindById(ctx context.Context, id int64) (*domain.Account, error){
	account, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return account, nil
}

func (s *AccountService) Transfer(ctx context.Context, fromID, toID int64, amount float64) error {
	err := s.repo.Transfer(ctx, fromID, toID, amount)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}