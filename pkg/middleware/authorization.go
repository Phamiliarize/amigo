package middleware

import (
	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"golang.org/x/exp/slices"
	"net/http"
)

type RouteTable map[string][]string

// Routes defines all reachable routes for the service; a route must be registered here to be exposed.
// TODO: How might we streamline this on the developer side
// TODO: Need to handle methods :)
var Routes = RouteTable{
	"/":           []string{"unauthenticated", "authenticated"},
	"/asd/{test}": []string{"authenticated"},
}

// Authorizer implements rudimentary RBAC and checks that a users role is correct for a requested route
func Authorizer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(dto.User)
		// TODO: We need to actually figure out how to run this post-routing :P
		// Since authorizer runs pre-routing, we need to ensure the requested route exists
		route, ok := Routes[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		for _, v := range route {
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
