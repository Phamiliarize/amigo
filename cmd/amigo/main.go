package main

import (
	"fmt"
	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/go-chi/chi/v5"
	"net/http"

	mw "github.com/Phamiliarize/amigo/pkg/middleware"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(mw.Authenticator)
	r.Use(mw.Authorizer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(dto.User)
		fmt.Printf("Body of handler: %v\n", user)
		w.Write([]byte("Hello World"))
	})

	r.Get("/asd/{test}", func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(dto.User)
		fmt.Printf("Body of handler: %v\n", user)
		w.Write([]byte("Hello World"))
	})

	http.ListenAndServe(":3000", r)
}
