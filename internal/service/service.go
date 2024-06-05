package service

import (
	"backend/internal/repo"
)

type Service interface {
}

type service struct {
	Repo repo.Repository
}

func NewService(r repo.Repository) Service {
	return &service{
		Repo: r,
	}
}
