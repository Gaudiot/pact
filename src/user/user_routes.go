package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	student_repository "github.com/pact/core/database/repository/student"
)

type userRouter struct {
	studentRepository student_repository.Repository
}

func NewUserRouter(repository student_repository.Repository) *userRouter {
	return &userRouter{studentRepository: repository}
}

func (router *userRouter) Router(r chi.Router) {
	r.Get("/id/{ID}", router.handleGetUser)
	r.Get("/", router.handleGetUsers)
	r.Post("/", router.handleCreateUser)
}

func (router *userRouter) handleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userID := chi.URLParam(r, "ID")

	user, err := router.studentRepository.GetByID(ctx, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (router *userRouter) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		http.Error(w, "Missing \"Name\" param", http.StatusBadRequest)
	}

	users, err := router.studentRepository.GetByName(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (router *userRouter) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" {
		http.Error(w, errors.New("").Error(), http.StatusBadRequest)
		return
	}

	userService := NewUserService(router.studentRepository)
	user, err := userService.CreateUser(req.Name, req.Email)
	if err != nil {
		http.Error(w, errors.New("").Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// MARK: - Improve later

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
