package rest

import (
	"github.com/go-chi/chi/v5"
)

func InitializeRoutes(r *chi.Mux) {
	r.Post("/file-upload/upload", FileUploadHandler)

}
