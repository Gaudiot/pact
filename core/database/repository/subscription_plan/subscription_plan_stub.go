package subscription_plan_repository

import (
	"context"
	"uuid"

	"github.com/pact/core/database/repository"
)

var _ Repository = &stub{}

type stub struct {
	subscriptionPlans map[string]SubscriptionPlan
}

func NewStub() *stub {
	return &stub{
		subscriptionPlans: map[string]SubscriptionPlan{
			"1": {
				Id:           "1",
				Title:        "Basic",
				Description:  "Basic subscription plan",
				PriceInCents: 2000,
			},
			"2": {
				Id:           "2",
				Title:        "Pro",
				Description:  "",
				PriceInCents: 5099,
			},
			"3": {
				Id:           "3",
				Title:        "Max",
				Description:  "",
				PriceInCents: 10000,
			},
		},
	}
}

func (s *stub) GetAll(ctx context.Context, pagination *repository.Pagination) (*repository.PaginationResponse[SubscriptionPlan], error) {
	p := []SubscriptionPlan{}
	for _, plan := range s.subscriptionPlans {
		p = append(p, plan)
	}

	totalPages := (len(s.subscriptionPlans) + pagination.Limit - 1) / pagination.Limit
	hasNextPage := totalPages > (pagination.Page + 1)

	start := pagination.Page * pagination.Limit
	if start >= len(p) {
		return &repository.PaginationResponse[SubscriptionPlan]{
			Data:        []SubscriptionPlan{},
			HasNextPage: false,
			TotalPages:  totalPages,
		}, nil
	}
	end := min(start+pagination.Limit, len(p))
	plans := p[start:end]
	return &repository.PaginationResponse[SubscriptionPlan]{
		Data:        plans,
		HasNextPage: hasNextPage,
		TotalPages:  totalPages,
	}, nil
}

func (s *stub) GetById(ctx context.Context, id string) (*SubscriptionPlan, error) {
	plan, ok := s.subscriptionPlans[id]
	if !ok {
		return nil, ErrSubscriptionPlanNotFound
	}

	return &plan, nil
}

func (s *stub) Create(ctx context.Context, params *CreateParams) (*SubscriptionPlan, error) {
	id := uuid.NewV4().String()
	plan := &SubscriptionPlan{
		Id:           id,
		Title:        params.Title,
		Description:  params.Description,
		PriceInCents: params.PriceInCents,
	}
	s.subscriptionPlans[id] = *plan

	return plan, nil
}
