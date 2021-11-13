package model

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID        int64          `gorm:"column:id;primary_key"  json:"id"`
	FullName  string         `gorm:"column:full_name"  json:"full_name"`
	BirthDay  int64          `gorm:"column:birth_day"  json:"birth_day"`
	CreatedAt time.Time      `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"  json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"  json:"deleted_at"`
}
