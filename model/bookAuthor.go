package model

type BookAuthor struct {
	ID       int64 `gorm:"column:id;primary_key"  json:"id"`
	BookID   int64 `gorm:"column:book_id"  json:"book_id"`
	AuthorID int64 `gorm:"column:author_id"  json:"author_id"`
}
