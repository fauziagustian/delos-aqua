package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId    int            `gorm:"primaryKey;column:user_id"`
	Name      string         `gorm:"column:name"`
	Email     string         `gorm:"column:email"`
	Password  string         `gorm:"column:password;not null" json:"-"`
	Token     string         `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
