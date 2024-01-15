package service

import (
	"context"
	"errors"

	"github.com/ivanldz/d-inventory/encryption"
	"github.com/ivanldz/d-inventory/internal/models"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("incorrect password")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {
	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u == nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)
	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.From64(u.Password)
	if err != nil {
		return nil, err
	}

	decrypterPassword, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(decrypterPassword) != password {
		return nil, ErrInvalidCredentials
	}

	return &models.User{
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
	}, nil
}
