package model

type BookRating struct {
	ID     int64 `gorm:"column:id;primary_key"  json:"id"`
	UserID int64 `gorm:"column:user_id"  json:"user_id"`
	BookID int64 `gorm:"column:book_id"  json:"book_id"`
	Rating int   `gorm:"column:rating"  json:"rating"`
}
