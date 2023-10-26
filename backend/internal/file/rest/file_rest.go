package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func FileStart() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.SetHeader("Access-Control-Allow-Credentials", "true"))

	InitializeRoutes(r)

	http.ListenAndServe(":8081", r)
}
