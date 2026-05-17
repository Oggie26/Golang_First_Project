package models

import (
	"time"

	"gorm.io/gorm"
)

type CategoryStatus string

const (
	CategoryStatusActive   CategoryStatus = "active"
	CategoryStatusInactive CategoryStatus = "inactive"
)

type Category struct {
	ID        int            `Gorm:"primaryKey;autoIncrement" db:"id" json:"id"`
	Name      string         `Gorm:"type:varchar(100);not null;uniqueIndex" db:"name" json:"name"`
	Status    CategoryStatus `Gorm:"type:text;not null;default:'active'" db:"status" json:"status"` // Dùng type:text để linh hoạt hơn trong migration
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt gorm.DeletedAt `Gorm:"index" json:"-" db:"deleted_at"`
}

func (Category) TableName() string {
	return "categories"
}

func init() {
	RegisterModel(&Category{})
}
