package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/kananniftiyev/cloudcourier-lib/shared"
)

func InitializeRoutes(r *chi.Mux) {
	r.Route("api/file", func(r chi.Router) {
		r.With(shared.JWTTokenVerifyMiddleware).Post("/upload", FileUploadHandler)
		r.With(shared.JWTTokenVerifyMiddleware).Get("/retrieve", FileRetrieveHandler)
		r.With(shared.JWTTokenVerifyMiddleware).Get("/history", FileUploadHistory)
	})
	

}
