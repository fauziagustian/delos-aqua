package models

import (
	"time"

	"gorm.io/gorm"
)

type Farm struct {
	FarmId    int            `gorm:"primaryKey;column:farm_id"`
	Name      string         `gorm:"column:name"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
