package repository

import (
	"GoLang_Tutorial/internal/models"
	"context"
)

type AccountRepository interface {
	Create(ctx context.Context, account *models.Account) error
	GetByUsername(ctx context.Context, username string) (*models.Account, error)
	Search(ctx context.Context, query string) ([]models.Account, error)
}
