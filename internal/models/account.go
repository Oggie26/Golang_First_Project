package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	RoleAdmin UserRole = "admin"
	RoleUser  UserRole = "user"
)

type Account struct {
	ID        uuid.UUID      `Gorm:"primaryKey;type:uuid" json:"id" db:"id"`
	Username  string         `Gorm:"uniqueIndex;not null" json:"username" db:"username"`
	Password  string         `Gorm:"not null" json:"-" db:"password"`
	Role      UserRole       `Gorm:"type:varchar(20);not null;default:'user'" json:"role" db:"role"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt gorm.DeletedAt `Gorm:"index" json:"-" db:"deleted_at"`

	User User `Gorm:"foreignKey:AccountID" json:"user"`
}

func (Account) TableName() string {
	return "accounts"
}

func init() {
	RegisterModel(&Account{})
}
