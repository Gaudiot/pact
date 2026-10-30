package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	student_repository "github.com/pact/core/database/repository/student"
	"github.com/pact/core/utils"
)

type userRouter struct {
	userService UserService
}

func NewUserRouter(userService *UserService) *userRouter {
	return &userRouter{userService: *userService}
}

func (router *userRouter) Router(r chi.Router) {
	r.Get("/id/{ID}", router.handleGetUser)
	r.Get("/", router.handleGetUsers)
	r.Post("/", router.handleCreateUser)
}

func (router *userRouter) handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "ID")

	user, err := router.userService.GetUser(r.Context(), userID)
	if err != nil {
		if errors.Is(err, student_repository.ErrStudentNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (router *userRouter) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	params := &student_repository.GetUsersParams{}

	if name := query.Get("name"); name != "" {
		params.Name = name
	}

	page, limit, err := utils.ParsePagination(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := router.userService.ListUsers(
		r.Context(),
		ListUserParams{
			Name:  params.Name,
			Page:  page,
			Limit: limit,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (router *userRouter) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Name == "" || req.Email == "" {
		http.Error(w, errors.New("").Error(), http.StatusBadRequest)
		return
	}

	user, err := router.userService.CreateUser(r.Context(), req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// MARK: - Improve later

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
