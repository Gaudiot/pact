package subscription_plan

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pact/core/utils"
)

type subscriptionPlanRouter struct {
	subscriptionPlanService SubscriptionPlanService
}

func NewSubscriptionPlanRouter(subscriptionPlanService *SubscriptionPlanService) *subscriptionPlanRouter {
	return &subscriptionPlanRouter{subscriptionPlanService: *subscriptionPlanService}
}

func (sr *subscriptionPlanRouter) Router(router chi.Router) {
	router.Get("/", sr.handleListPlans)
	router.Post("/", sr.handleCreate)
}

func (sr *subscriptionPlanRouter) handleListPlans(w http.ResponseWriter, r *http.Request) {
	page, limit, err := utils.ParsePagination(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	plans, err := sr.subscriptionPlanService.ListPlans(r.Context(), ListPlansParams{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(plans)
}

func (sr *subscriptionPlanRouter) handleCreate(w http.ResponseWriter, r *http.Request) {
	var req createSubscriptionPlanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.PriceInCents <= 0 {
		http.Error(w, errors.New("").Error(), http.StatusBadRequest)
		return
	}

	plan, err := sr.subscriptionPlanService.Create(r.Context(), CreateParams{Title: req.Title, Description: req.Description, PriceInCents: req.PriceInCents})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(plan)
}

// MARK: - Improve later

type createSubscriptionPlanRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	PriceInCents int    `json:"price_in_cents"`
}
