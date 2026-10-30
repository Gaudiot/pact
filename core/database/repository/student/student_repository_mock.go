package student_repository

import (
	"context"

	"github.com/pact/core/database/repository"
	"github.com/stretchr/testify/mock"
)

var _ Repository = &MockRepository{}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Create(ctx context.Context, params *CreateStudentParams) (*Student, error) {
	args := m.Called(ctx, params)

	response, _ := args.Get(0).(*Student)

	return response, args.Error(1)
}

func (m *MockRepository) GetByID(ctx context.Context, id string) (*Student, error) {
	args := m.Called(ctx, id)

	response, _ := args.Get(0).(*Student)

	return response, args.Error(1)
}

func (m *MockRepository) GetUsers(ctx context.Context, params *GetUsersParams, pagination *repository.Pagination) ([]*Student, error) {
	args := m.Called(ctx, params, pagination)

	response, _ := args.Get(0).([]*Student)

	return response, args.Error(1)
}
