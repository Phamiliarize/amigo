package middleware

import (
	"context"
	"net/http"

	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/google/uuid"
)

// Authenticator is a middleware that checks for the presence of a token, validates it, and injects the user
func Authenticator(next http.Handler) http.Handler {
	// TODO: Allow the authenticator to take input on where to pull JWK, and what the userID/roles keys in the token are.
	// TODO: actual user handling stoofs
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", dto.User{ID: uuid.NewString(), Roles: []string{"authenticated"}})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
