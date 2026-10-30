package student_repository

import (
	"context"
	"errors"

	"github.com/pact/core/database/repository"
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

type GetUsersParams struct {
	Name string
}

var (
	ErrStudentNotFound        = errors.New("Student not found")
	ErrEmailAlreadyRegistered = errors.New("Email already registered")
)

type Repository interface {
	Create(ctx context.Context, params *CreateStudentParams) (*Student, error)
	GetByID(ctx context.Context, id string) (*Student, error)
	GetUsers(ctx context.Context, params *GetUsersParams, pagination *repository.Pagination) ([]*Student, error)
}
