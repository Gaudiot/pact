package student_repository

import (
	"context"

	"github.com/stretchr/testify/mock"
)

var _ StudentRepository = &MockStudentRepository{}

type MockStudentRepository struct {
	mock.Mock
}

func (m *MockStudentRepository) Create(ctx context.Context, student *Student) error {
	args := m.Called(ctx, student)

	return args.Error(0)
}

func (m *MockStudentRepository) GetByID(ctx context.Context, id string) (*Student, error) {
	args := m.Called(ctx, id)

	response, _ := args.Get(0).(*Student)

	return response, args.Error(1)
}
