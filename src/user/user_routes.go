package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRouter(r chi.Router) {
	r.Get("/{ID}", handleGetUser)
	r.Post("/", handleCreateUser)
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "ID")
	var _ = userID
	w.WriteHeader(http.StatusTeapot)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusTeapot)
}

// MARK: - Improve later

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
