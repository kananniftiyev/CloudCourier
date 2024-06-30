package rest

import (
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {
	r.Post("/file/upload", FileUploadHandler)
	r.Get("/file/retrieve", FileRetrieveHandler)
	r.Get("/file/history", FileUploadHistory)

}
