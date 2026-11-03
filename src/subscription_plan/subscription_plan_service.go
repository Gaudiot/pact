package subscription_plan

import (
	"context"

	"github.com/pact/core/database/repository"
	subscription_plan_repository "github.com/pact/core/database/repository/subscription_plan"
)

type SubscriptionPlan struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	PriceInCents int    `json:"price_in_cents"`
}

type SubscriptionPlanService struct {
	subscriptionPlanRepository subscription_plan_repository.Repository
}

func NewSubscriptionPlanService(subscriptionPlanRepository subscription_plan_repository.Repository) *SubscriptionPlanService {
	return &SubscriptionPlanService{subscriptionPlanRepository: subscriptionPlanRepository}
}

type ListPlansParams struct {
	Page  int
	Limit int
}

func (s *SubscriptionPlanService) ListPlans(ctx context.Context, params ListPlansParams) ([]SubscriptionPlan, error) {
	result, err := s.subscriptionPlanRepository.GetAll(ctx, &repository.Pagination{Page: params.Page, Limit: params.Limit})
	if err != nil {
		return nil, err
	}

	plans := []SubscriptionPlan{}
	for _, plan := range result.Data {
		plans = append(plans, SubscriptionPlan{
			Id:           plan.Id,
			Title:        plan.Title,
			Description:  plan.Description,
			PriceInCents: plan.PriceInCents,
		})
	}
	return plans, nil
}

type CreateParams struct {
	Title        string
	Description  string
	PriceInCents int
}

func (s *SubscriptionPlanService) Create(ctx context.Context, params CreateParams) (*SubscriptionPlan, error) {
	result, err := s.subscriptionPlanRepository.Create(ctx, &subscription_plan_repository.CreateParams{Title: params.Title, Description: params.Description, PriceInCents: params.PriceInCents})
	if err != nil {
		return nil, err
	}

	return &SubscriptionPlan{
		Id:           result.Id,
		Title:        params.Title,
		Description:  params.Description,
		PriceInCents: params.PriceInCents,
	}, nil
}
