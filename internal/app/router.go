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

	return http.ListenAndServe(":8080", nil)
}
