package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid" json:"id" db:"id"`
	AccountID uuid.UUID      `gorm:"uniqueIndex;not null;type:uuid" json:"account_id" db:"account_id"`
	FullName  string         `gorm:"type:varchar(100)" json:"full_name" db:"full_name"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex" json:"email" db:"email"`
	Avatar    string         `json:"avatar" db:"avatar"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-" db:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}

func init() {
	RegisterModel(&User{})
}
