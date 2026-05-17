package repository

import (
	"GoLang_Tutorial/internal/models"
	"GoLang_Tutorial/pkg/database"

	"context"
)

type accountRepositoryPostgres struct {
	db *database.Database
}

func NewAccountRepository(db *database.Database) AccountRepository {
	return &accountRepositoryPostgres{db: db}
}

func (r *accountRepositoryPostgres) GetDB() *database.Database {
	return r.db
}

func (r *accountRepositoryPostgres) Create(ctx context.Context, account *models.Account) error {
	return r.db.Gorm.WithContext(ctx).Create(account).Error
}

func (r *accountRepositoryPostgres) GetByUsername(ctx context.Context, username string) (*models.Account, error) {
	sql := `SELECT id, username, password, created_at, updated_at FROM accounts 
			WHERE username = $1 AND deleted_at IS NULL`

	var account models.Account

	err := r.db.Sqlx.GetContext(ctx, &account, sql, username)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepositoryPostgres) Search(ctx context.Context, query string) ([]models.Account, error) {
	sql := `SELECT id, username, created_at FROM accounts 
			WHERE username ILIKE $1 AND deleted_at IS NULL 
			ORDER BY created_at DESC`

	var accounts []models.Account
	err := r.db.Sqlx.SelectContext(ctx, &accounts, sql, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
