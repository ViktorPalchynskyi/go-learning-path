package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/username/lesson_13/internal/domain"
)

type AccountRepository struct {
	pool *pgxpool.Pool
}

func NewAccountRepository(pool *pgxpool.Pool) *AccountRepository{
	return &AccountRepository{pool: pool}
}

func (r *AccountRepository) Create(ctx context.Context, ownerID int64, balance float64) (*domain.Account, error){
	var (
		account domain.Account
		pgErr *pgconn.PgError
	)
	err := r.pool.QueryRow(ctx,
	"INSERT INTO accounts (owner_id, balance) VALUES ($1, $2) RETURNING id, balance, owner_id",
	ownerID, balance,
	).Scan(&account.ID, &account.Balance, &account.OwnerId)

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return nil, domain.ErrAlreadyExists
		}
	}

	return &account, nil
}

func (r *AccountRepository) FindById(ctx context.Context, id int64) (*domain.Account, error){
	var account domain.Account
	err := r.pool.QueryRow(ctx,
	"SELECT id, balance, owner_id FROM accounts WHERE id = $1",
	id,
	).Scan(&account.ID, &account.Balance, &account.OwnerId)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNotFound
	}

	return &account, nil
}

func (r *AccountRepository) Transfer(ctx context.Context, fromID, toID int64, amount float64) error {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, 
		"UPDATE accounts SET balance = balance - $1 WHERE id = $2", 
		amount, fromID,
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, 
		"UPDATE accounts SET balance = balance + $1 WHERE id = $2", 
		amount, toID,
	)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}