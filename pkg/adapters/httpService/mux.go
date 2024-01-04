package httpService

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Phamiliarize/amigo/pkg/adapters/httpService/api"
	amigoMW "github.com/Phamiliarize/amigo/pkg/adapters/httpService/middleware"
	"github.com/Phamiliarize/amigo/pkg/adapters/httpService/views"
	"github.com/Phamiliarize/amigo/pkg/application/port"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Routes defines all reachable routes for the service from an authorization standpoint;
var routes = amigoMW.RouteTable{}

// route function registers a route adding it to the MUX and the Route Table
func route(r chi.Router, method string, path string, roles []string, handler http.HandlerFunc) {
	r.MethodFunc(method, path, handler)
	routes[path] = map[string][]string{}
	routes[path][method] = roles
}

// NewAmigoHTTPServer initializes an instance of the Amigo HTTP Server/MUX
func NewAmigoHTTPServer(jsonAPI api.JsonAPI, viewCollection views.ViewCollection, ThemeService port.ThemeService) *chi.Mux {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(middleware.RedirectSlashes)
	r.Use(amigoMW.Authenticator)

	// Initialize the authorizer middleware
	authorizer := amigoMW.NewAuthorizerMiddleware(&routes)

	//Expose routes for themes public assets folders
	workDir, _ := os.Getwd()
	for _, theme := range ThemeService.GetThemes() {
		if theme.Publish {
			themeAssetsPath := fmt.Sprintf("%s/assets", theme.Path)[1:]
			filesDir := http.Dir(filepath.Join(workDir, themeAssetsPath))
			route := fmt.Sprintf("/assets/%s", theme.Dir)
			FileServer(r, route, filesDir)
		}
	}

	// Routes
	r.Group(func(r chi.Router) {
		// Authorizer needs to be beneath the group in order to receive the routed path
		r.Use(authorizer.Authorizer)

		// JSON API Routes
		route(r, "GET", "/api/me", []string{"authenticated"}, jsonAPI.GetMe)
		route(r, "PATCH", "/api/preference", []string{"authenticated"}, jsonAPI.PatchPreferences)

		// Views
		route(r, "GET", "/", []string{"unauthenticated", "authenticated"}, viewCollection.Home)
	})

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
