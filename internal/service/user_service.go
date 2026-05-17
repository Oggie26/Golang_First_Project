package service

import (
	"context"

	"GoLang_Tutorial/internal/dto"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
}
