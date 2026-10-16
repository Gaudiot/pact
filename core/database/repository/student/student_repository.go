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
	ErrStudentNotFound = errors.New("Student not found")
)

type Repository interface {
	Create(ctx context.Context, params *CreateStudentParams) (*Student, error)
	GetByID(ctx context.Context, id string) (*Student, error)
}
