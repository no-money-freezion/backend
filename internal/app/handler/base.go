package handler

import (
	"backend/internal/repo"
	shortener "backend/internal/service"
)

type Handler struct {
	Service shortener.Service
	Repo    repo.Repository
}

func NewHandler(s shortener.Service, repo repo.Repository) *Handler {
	return &Handler{
		Service: s,
		Repo:    repo,
	}
}
