package middleware

import (
	"net/http"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

const (
	ErrMissingOrInvalidBearerToken string = "missing or invalid bearer token"
	ErrInvalidOrExpiredToken       string = "invalid or expired token"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return clerkhttp.RequireHeaderAuthorization()(next)
}
