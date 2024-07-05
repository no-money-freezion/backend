// Package app Description: This file contains the router logic for the application.
package app

import (
	"backend/internal/app/handler"
	"backend/internal/repo"
	"backend/internal/service"
	"github.com/rs/cors"
	"net/http"
)

type App struct {
	Service service.Service
	Repo    repo.Repository
}

func NewApp(s service.Service, r repo.Repository) *App {
	return &App{
		Service: s,
		Repo:    r,
	}
}

func (r *App) Start() error {
	h := handler.NewHandler(r.Service, r.Repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.HealthHandler)
	mux.HandleFunc("/auth/register", h.RegisterHandler)
	mux.HandleFunc("/auth/login", h.LoginHandler)
	mux.HandleFunc("/auth/password-recovery", h.PasswordRecoveryHandler)
	mux.HandleFunc("/auth/password-reset", h.PasswordResetHandler)
	mux.HandleFunc("/roulette/info", h.RouletteInfoHandler)
	mux.HandleFunc("/roulette/bet", h.RouletteBetHandler)

	// Define CORS options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://example.com", "http://anotherdomain.com", "*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap the mux with the CORS middleware
	handler := c.Handler(mux)

	return http.ListenAndServe(":8080", handler)
}
