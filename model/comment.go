package model

import (
	"time"

	bookmodel "github.com/cnson19700/book_service/model"
	"gorm.io/gorm"
)

type Comment struct {
	ID        int64          `gorm:"column:id;primary_key"  json:"id"`
	ParentID  int64          `gorm:"column:parent_id; default:1"  json:"parent_id"`
	UserID    int64          `gorm:"column:user_id" json:"user_id"`
	BookID    int64          `gorm:"column:book_id" json:"book_id"`
	Content   string         `gorm:"column:content" json:"content"`
	CreatedAt time.Time      `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at"  json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"  json:"deleted_at"`
}

type CommentFilter struct {
	BookID   int64 `json:"book_id"`
	ParentID int64 `json:"parent_id"`
}

type CommentResult struct {
	Data []Comment `json:"data"`
	bookmodel.Paginator
}
