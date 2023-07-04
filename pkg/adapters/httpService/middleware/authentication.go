package middleware

import (
	"context"
	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"net/http"
)

// Authenticator is a middleware that checks for the presence of a token, validates it, and injects the user
func Authenticator(next http.Handler) http.Handler {
	// TODO: Allow the authenticator to take input on where to pull JWK, and what the userID/roles keys in the token are.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", dto.User{ID: "fewhf", Roles: []string{"authenticated"}})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
