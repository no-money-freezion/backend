// Package app Description: This file contains the router logic for the application.
package app

import (
	"backend/internal/app/handler"
	"backend/internal/repo"
	"backend/internal/service"
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

	http.HandleFunc("GET /health", h.HealthHandler)
	http.HandleFunc("POST /auth/register", h.RegisterHandler)
	http.HandleFunc("POST /auth/login", h.LoginHandler)
	http.HandleFunc("POST /auth/password-recovery", h.PasswordRecoveryHandler)
	http.HandleFunc("POST /auth/password-reset", h.PasswordResetHandler)
	http.HandleFunc("POST /roulette/info", h.RouletteInfoHandler)
	http.HandleFunc("POST /roulette/bet", h.RouletteBetHandler)

	return http.ListenAndServe(":8080", nil)
}
