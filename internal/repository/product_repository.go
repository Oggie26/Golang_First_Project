package repository

import (
	"GoLang_Tutorial/internal/models"
	"context"
)

type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id uint) (*models.Product, error)
	ListByCategory(ctx context.Context, categoryID uint) ([]models.Product, error)
}
