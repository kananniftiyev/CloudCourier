package rest

import "github.com/go-chi/chi/v5"

func InitializeRoutes(r *chi.Mux) {
	r.Post("/auth/register", RegisterHandler)
	r.Post("/auth/login", LoginHandler)
}
