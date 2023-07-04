package middleware

import (
	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slices"
	"net/http"
)

// RouteTable is meant to hold the mapping of routes and their allowed RBAC roles
type RouteTable map[string]map[string][]string

type authorizerMiddleware struct {
	RouteTable *RouteTable
}

// NewAuthorizerMiddleware returns an Authorizer Middleware
func NewAuthorizerMiddleware(routeTable *RouteTable) authorizerMiddleware {
	return authorizerMiddleware{
		RouteTable: routeTable,
	}
}

// Authorizer implements rudimentary RBAC and checks that a users role is correct for a requested route
func (a *authorizerMiddleware) Authorizer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		routes := *a.RouteTable
		user := r.Context().Value("user").(dto.User)
		routePattern := chi.RouteContext(r.Context()).RoutePattern()

		// Handle root being treated as empty string
		if routePattern == "" {
			routePattern = "/"
		}

		// Any route not found in the route table should be treated as a 404
		route, ok := routes[routePattern]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// Any method found for a route is treated as a 404 too
		roles, ok := route[r.Method]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		for _, v := range roles {
			// If the role is a match, continue silently. We might consider appending the winning role to context
			if slices.Contains(user.Roles, v) {
				next.ServeHTTP(w, r)
			}
		}

		// Default for this function is to return forbidden
		w.WriteHeader(http.StatusForbidden)
		return
	})
}
