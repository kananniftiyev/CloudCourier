package rest

import (
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {
	r.Route("api/file", func(r chi.Router) {
		r.With(JWTTokenVerifyMiddleware).Post("/upload", FileUploadHandler)
		r.With(JWTTokenVerifyMiddleware).Get("/retrieve", FileRetrieveHandler)
		r.With(JWTTokenVerifyMiddleware).Get("/history", FileUploadHistory)
	})
	

}
