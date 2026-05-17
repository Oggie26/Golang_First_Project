package repository

import (
	"GoLang_Tutorial/internal/models"
	"GoLang_Tutorial/pkg/database"
	"context"
)

type productRepositoryPostgres struct {
	db *database.Database
}

func NewProductRepository(db *database.Database) ProductRepository {
	return &productRepositoryPostgres{db: db}
}

func (r *productRepositoryPostgres) Create(ctx context.Context, product *models.Product) error {
	return r.db.Gorm.WithContext(ctx).Create(product).Error
}

func (r *productRepositoryPostgres) GetByID(ctx context.Context, id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Gorm.WithContext(ctx).Preload("Category").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepositoryPostgres) ListByCategory(ctx context.Context, categoryID uint) ([]models.Product, error) {
	// Dùng sqlx: Khi cần Join phức tạp hoặc lấy danh sách lớn để tối ưu tốc độ
	sql := `SELECT p.*, c.name as "category.name" 
			FROM products p 
			LEFT JOIN categories c ON p.category_id = c.id 
			WHERE p.category_id = $1 AND p.deleted_at IS NULL`

	var products []models.Product
	err := r.db.Sqlx.SelectContext(ctx, &products, sql, categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}
