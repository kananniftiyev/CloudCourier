package rest

import (
	auth "backend/auth-service/internal"

	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {
	r.Post("/auth/register", RegisterHandler)
	r.Post("/auth/login", LoginHandler)
	r.With(auth.JWTTokenVerifyMiddleware).Post("/auth/logout", LogoutHandler)
	r.With(auth.JWTTokenVerifyMiddleware).Get("/auth/user", UserHandler)
}
