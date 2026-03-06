package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/username/lesson_13/internal/domain"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository{
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, name, email string) (*domain.User, error){
	var (
		user domain.User
		pgErr *pgconn.PgError
	)
	err := r.pool.QueryRow(ctx,
	"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id, name, email",
	name, email,
	).Scan(&user.ID, &user.Name, &user.Email)

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return nil, domain.ErrAlreadyExists
		}
	}

	return &user, nil
}

func (r *UserRepository) FindById(ctx context.Context, id int64) (*domain.User, error){
	var user domain.User
	err := r.pool.QueryRow(ctx,
	"SELECT id, name, email FROM users WHERE id = $1",
	id,
	).Scan(&user.ID, &user.Name, &user.Email)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrNotFound
	}

	return &user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*domain.User, error){
	var users []*domain.User
	rows, err := r.pool.Query(ctx,
	"SELECT id, name, email FROM users",
	)
	if err != nil {                                                                                                                                                                             
		return nil, err
	} 
	defer rows.Close()
	for rows.Next() {
		var user domain.User
		rows.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, &user)
	}

	return users, nil
}

func (r *UserRepository) Delete(ctx context.Context, id int64) error{
	_, err := r.pool.Exec(ctx, "DELETE FROM users WHERE id = $1", id)

	if errors.Is(err, pgx.ErrNoRows) {
		return domain.ErrNotFound
	}

	return err
}