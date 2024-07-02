package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/kananniftiyev/cloudcourier-lib/shared"
)

func InitializeRoutes(r *chi.Mux) {
	r.Route("api/auth", func(r chi.Router) {
		r.Post("/register", RegisterHandler)
		r.Post("/login", LoginHandler)
		r.With(shared.JWTTokenVerifyMiddleware).Post("/logout", LogoutHandler)
		r.With(shared.JWTTokenVerifyMiddleware).Get("/user", UserHandler)
	})
}
