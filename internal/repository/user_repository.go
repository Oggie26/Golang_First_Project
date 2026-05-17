package repository

import (
	"context"
	"GoLang_Tutorial/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
}
