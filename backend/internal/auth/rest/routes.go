package rest

import (
	"backend/internal/auth"
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {
	r.Post("/auth/register", RegisterHandler)
	r.Post("/auth/login", LoginHandler)
	r.With(auth.VerifyToken).Post("/auth/logout", LogoutHandler)
	r.With(auth.VerifyToken).Get("/auth/user", UserHandler)
}
