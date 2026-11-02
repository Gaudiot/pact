package subscription_plan_repository

import (
	"context"
	"errors"

	"github.com/pact/core/database/repository"
)

type SubscriptionPlan struct {
	Id           string
	Title        string
	Description  string
	PriceInCents int
}

type CreateParams struct {
	Title        string
	Description  string
	PriceInCents int
}

var (
	ErrSubscriptionPlanNotFound = errors.New("Subscription plan not found")
)

type Repository interface {
	GetById(ctx context.Context, id string) (*SubscriptionPlan, error)
	GetAll(ctx context.Context, pagination *repository.Pagination) ([]SubscriptionPlan, error)
	Create(ctx context.Context, params *CreateParams) (*SubscriptionPlan, error)
}
