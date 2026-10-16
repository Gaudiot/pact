package student_repository

import "context"

type Student struct {
	ID    string
	Name  string
	Email string
}

type StudentRepository interface {
	Create(ctx context.Context, student *Student) error
	GetByID(ctx context.Context, id string) (*Student, error)
}
