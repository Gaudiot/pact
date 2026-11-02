package subscription_plan_repository

import (
	"context"

	"github.com/pact/core/database/repository"
	"github.com/stretchr/testify/mock"
)

var _ Repository = &MockRepository{}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAll(ctx context.Context, pagination *repository.Pagination) ([]SubscriptionPlan, error) {
	args := m.Called(ctx, pagination)

	response, _ := args.Get(0).([]SubscriptionPlan)

	return response, args.Error(1)
}

func (m *MockRepository) GetById(ctx context.Context, id string) (*SubscriptionPlan, error) {
	args := m.Called(ctx, id)

	response, _ := args.Get(0).(*SubscriptionPlan)

	return response, args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, params *CreateParams) (*SubscriptionPlan, error) {
	args := m.Called(ctx, params)

	response, _ := args.Get(0).(*SubscriptionPlan)

	return response, args.Error(1)
}
