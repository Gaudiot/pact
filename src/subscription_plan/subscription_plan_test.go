package subscription_plan

import (
	"errors"
	"testing"

	"github.com/pact/core/database/repository"
	subscription_plan_repository "github.com/pact/core/database/repository/subscription_plan"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupTest(t *testing.T) (
	*SubscriptionPlanService,
	*subscription_plan_repository.MockRepository,
) {
	t.Helper()

	mockSubscriptionPlanRepository := new(subscription_plan_repository.MockRepository)
	sut := NewSubscriptionPlanService(mockSubscriptionPlanRepository)

	return sut, mockSubscriptionPlanRepository
}

func TestListPlans(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, mockSubscriptionPlanRepository := setupTest(t)
		ctx := t.Context()
		params := ListPlansParams{Page: 2, Limit: 10}
		expected := []SubscriptionPlan{
			{Id: "basic-id", Title: "Basic", Description: "Basic plan", PriceInCents: 1990},
			{Id: "premium-id", Title: "Premium", Description: "Premium plan", PriceInCents: 4990},
		}
		mockSubscriptionPlanRepository.On("GetAll", ctx,
			&repository.Pagination{Page: params.Page, Limit: params.Limit},
		).Return([]subscription_plan_repository.SubscriptionPlan{
			{Id: "basic-id", Title: "Basic", Description: "Basic plan", PriceInCents: 1990},
			{Id: "premium-id", Title: "Premium", Description: "Premium plan", PriceInCents: 4990},
		}, nil).Once()

		plans, err := sut.ListPlans(ctx, params)

		require.NoError(t, err)
		assert.Equal(t, expected, plans)
		mockSubscriptionPlanRepository.AssertExpectations(t)
	})
	t.Run("No Plans Found", func(t *testing.T) {
		tests := []struct {
			testName string
			plans    []subscription_plan_repository.SubscriptionPlan
		}{
			{"Empty Result", []subscription_plan_repository.SubscriptionPlan{}},
			{"Nil Result", nil},
		}

		for _, tt := range tests {
			t.Run(tt.testName, func(t *testing.T) {
				sut, mockSubscriptionPlanRepository := setupTest(t)
				ctx := t.Context()
				params := ListPlansParams{Limit: 10}
				mockSubscriptionPlanRepository.On("GetAll", ctx,
					&repository.Pagination{Page: params.Page, Limit: params.Limit},
				).Return(tt.plans, nil).Once()

				plans, err := sut.ListPlans(ctx, params)

				require.NoError(t, err)
				assert.NotNil(t, plans)
				assert.Empty(t, plans)
				mockSubscriptionPlanRepository.AssertExpectations(t)
			})
		}
	})
	t.Run("Repository Error", func(t *testing.T) {
		sut, mockSubscriptionPlanRepository := setupTest(t)
		ctx := t.Context()
		params := ListPlansParams{Limit: 10}
		expectedError := errors.New("internal_error")
		mockSubscriptionPlanRepository.On("GetAll", ctx,
			&repository.Pagination{Page: params.Page, Limit: params.Limit},
		).Return(nil, expectedError).Once()

		plans, err := sut.ListPlans(ctx, params)

		require.ErrorIs(t, err, expectedError)
		require.Nil(t, plans)
		mockSubscriptionPlanRepository.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		sut, mockSubscriptionPlanRepository := setupTest(t)
		ctx := t.Context()
		params := CreateParams{Title: "Basic", Description: "Basic plan", PriceInCents: 1990}
		expected := &SubscriptionPlan{
			Id: "basic-id", Title: params.Title, Description: params.Description, PriceInCents: params.PriceInCents,
		}
		mockSubscriptionPlanRepository.On("Create", ctx,
			&subscription_plan_repository.CreateParams{
				Title: params.Title, Description: params.Description, PriceInCents: params.PriceInCents,
			},
		).Return(&subscription_plan_repository.SubscriptionPlan{
			Id: expected.Id, Title: params.Title, Description: params.Description, PriceInCents: params.PriceInCents,
		}, nil).Once()

		plan, err := sut.Create(ctx, params)

		require.NoError(t, err)
		require.NotNil(t, plan)
		assert.Equal(t, expected, plan)
		mockSubscriptionPlanRepository.AssertExpectations(t)
	})
	t.Run("Repository Error", func(t *testing.T) {
		sut, mockSubscriptionPlanRepository := setupTest(t)
		ctx := t.Context()
		params := CreateParams{Title: "Basic", Description: "Basic plan", PriceInCents: 1990}
		expectedError := errors.New("internal_error")
		mockSubscriptionPlanRepository.On("Create", ctx,
			&subscription_plan_repository.CreateParams{
				Title: params.Title, Description: params.Description, PriceInCents: params.PriceInCents,
			},
		).Return(nil, expectedError).Once()

		plan, err := sut.Create(ctx, params)

		require.ErrorIs(t, err, expectedError)
		require.Nil(t, plan)
		mockSubscriptionPlanRepository.AssertExpectations(t)
	})
}
