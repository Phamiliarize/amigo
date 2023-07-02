package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	mw "github.com/Phamiliarize/amigo/pkg/middleware"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(mw.MyMiddleware)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(string)
		w.Write([]byte(user))
	})
	http.ListenAndServe(":3000", r)
}
