package rest

import (
	auth "backend/auth-service/internal"

	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {
	r.Route("api/auth", func(r chi.Router) {
		r.Post("/register", RegisterHandler)
		r.Post("/login", LoginHandler)
		r.With(auth.JWTTokenVerifyMiddleware).Post("/logout", LogoutHandler)
		r.With(auth.JWTTokenVerifyMiddleware).Get("/user", UserHandler)
	})
}
