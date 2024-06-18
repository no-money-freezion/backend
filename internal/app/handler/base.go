package handler

import (
	"backend/internal/repo"
	shortener "backend/internal/service"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Service   shortener.Service
	Repo      repo.Repository
	Validator *validator.Validate
}

func NewHandler(s shortener.Service, repo repo.Repository) *Handler {
	return &Handler{
		Service:   s,
		Repo:      repo,
		Validator: validator.New(),
	}
}

func (h *Handler) Validate(i interface{}) error {
	return h.Validator.Struct(i)
}
