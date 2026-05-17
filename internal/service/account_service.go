package service

import (
	"GoLang_Tutorial/internal/dto"
	"GoLang_Tutorial/internal/models"
	"context"
)

type AccountService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	GetProfile(ctx context.Context, accountID string) (*models.User, error)
	UpdateProfile(ctx context.Context, accountID string, fullName, email string) error
}
