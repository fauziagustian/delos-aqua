package models

import (
	"time"

	"gorm.io/gorm"
)

type Pond struct {
	PondId    int            `gorm:"primaryKey;column:pond_id"`
	Name      string         `gorm:"column:name"`
	FarmID    string         `gorm:"foreignKey:FarmId;column:farm_id"`
	CreatedAt time.Time      `gorm:"column:created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}
