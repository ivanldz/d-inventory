package service

import (
	"context"

	"github.com/ivanldz/d-inventory/internal/models"
	"github.com/ivanldz/d-inventory/internal/repository"
)

type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
