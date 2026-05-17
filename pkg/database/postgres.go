package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Gorm *gorm.DB
	Sqlx *sqlx.DB
}

func NewPostgresConnection(url string) (*Database, error) {
	// 1. Kết nối GORM
	gormDB, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 2. Lấy SQL DB gốc
	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	// 3. Bọc bằng sqlx
	sqlxDB := sqlx.NewDb(sqlDB, "postgres")

	return &Database{
		Gorm: gormDB,
		Sqlx: sqlxDB,
	}, nil
}

// MigrateAll giúp bạn tự động migrate tất cả các models được truyền vào
func (db *Database) MigrateAll(models ...interface{}) error {
	log.Println("🚀 Đang chạy Database Migration...")
	return db.Gorm.AutoMigrate(models...)
}

func (db *Database) Close() {
	sqlDB, _ := db.Gorm.DB()
	sqlDB.Close()
}
