package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `Gorm:"primaryKey;autoIncrement" db:"id" json:"id"`
	Name        string         `Gorm:"type:varchar(255);not null" db:"name" json:"name"`
	Code        string         `Gorm:"type:varchar(50);uniqueIndex;not null" db:"code" json:"code"`
	Price       float64        `Gorm:"type:decimal(10,2);not null" db:"price" json:"price"`
	Stock       int            `Gorm:"not null;default:0" db:"stock" json:"stock"`
	Description string         `Gorm:"type:text" db:"description" json:"description"`
	CategoryID  uint           `Gorm:"index" db:"category_id" json:"category_id"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt   gorm.DeletedAt `Gorm:"index" json:"-" db:"deleted_at"`

	Category *Category `gorm:"foreignKey:CategoryID" json:"category"`
}

func (Product) TableName() string {
	return "products"
}

func init() {
	RegisterModel(&Product{})
}
