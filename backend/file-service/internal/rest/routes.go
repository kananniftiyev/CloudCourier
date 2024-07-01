package rest

import (
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {
	r.With(JWTTokenVerifyMiddleware).Post("/file/upload", FileUploadHandler)
	r.With(JWTTokenVerifyMiddleware).Get("/file/retrieve", FileRetrieveHandler)
	r.With(JWTTokenVerifyMiddleware).Get("/file/history", FileUploadHistory)

}
