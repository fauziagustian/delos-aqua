package models

import (
	"time"

	"gorm.io/gorm"
)

type UserAgents struct {
	UserAgentId int            `gorm:"primaryKey;column:user_id"`
	MethodUrl   string         `gorm:"column:method_url"`
	UserId      int            `gorm:"foreignKey:UserId;column:user_id"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
}
