package user

import (
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
	w.WriteHeader(http.StatusTeapot)
}
