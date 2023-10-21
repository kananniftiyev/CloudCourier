package rest

import (
	"backend/internal/auth"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func InitializeRoutes(r *chi.Mux) {
	r.Post("/auth/register", RegisterHandler)
	r.Post("/auth/login", LoginHandler)
	r.With(auth.VerifyToken).Get("/auth/upload", func(w http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("asdas"))
	})
}
