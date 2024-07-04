package rest

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kananniftiyev/cloudcourier-lib/shared"
	"golang.org/x/net/http2"
)

// TODO: Rate Limit
func AuthStart() {
	shared.LoadEnv()

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.CleanPath)
	// Debug || Profiler
	router.Mount("/debug", middleware.Profiler())


	InitializeRoutes(router)

	
	// Load TLS certificates
	cert, err := tls.LoadX509KeyPair(os.Getenv("PEM_FILE_PATH"), os.Getenv("KEY_FILE_PATH"))
	if err != nil {
		log.Fatalf("Error loading certificate: %v", err)
	}

	// TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// HTTP server configuration
	server := &http.Server{
		Addr:      ":8081",
		Handler:   router,
		TLSConfig: tlsConfig,
	}

	// Enable HTTP/2
	http2.ConfigureServer(server, &http2.Server{})

	// Log the server start information
	log.Printf("Starting HTTP/2 server on %s", server.Addr)

	// Start the HTTP/2 server
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatalf("Error starting HTTP/2 server: %v", err)
	}
	
}
