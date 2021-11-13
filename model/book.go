package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID            int64          `gorm:"column:id;primary_key"  json:"id"`
	AuthorID      int64          `gorm:"column:author_id" json:"author_id"`
	PublisherID   int64          `gorm:"column:publisher_id" json:"publisher_id"`
	Title         string         `gorm:"column:title" json:"title"`
	Language      string         `gorm:"column:language" json:"language"`
	Image         string         `gorm:"column:image" json:"image"`
	Description   string         `gorm:"column:description" json:"description"`
	ReleaseDate   time.Time      `gorm:"column:release_date" json:"release_date"`
	RatingAverage float64        `gorm:"column:rating_average" json:"rating_average"`
	TotalPage     int64          `gorm:"column:total_page" json:"total_page"`
	CreatedAt     time.Time      `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at"  json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at"  json:"deleted_at"`
}

type BookFilter struct {
	AuthorID  int64 `json:"author_id"`
	CateID    int64 `json:"cate_id"`
	MinRating int   `json:"min_rating"` //default value no filter: -1
}

type BookResult struct {
	Data []Book `json:"data"`
	Paginator
}
