package student_repository

import (
	"context"
	"errors"
)

type Student struct {
	ID    string
	Name  string
	Email string
}

type CreateStudentParams struct {
	Name  string
	Email string
}

var (
	ErrStudentNotFound        = errors.New("Student not found")
	ErrEmailAlreadyRegistered = errors.New("Email already registered")
)

type Repository interface {
	Create(ctx context.Context, params *CreateStudentParams) (*Student, error)
	GetByID(ctx context.Context, id string) (*Student, error)
	GetByName(ctx context.Context, name string) ([]*Student, error)
}
