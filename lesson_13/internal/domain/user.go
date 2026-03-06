package domain

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}