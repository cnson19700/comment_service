package model

type BookCategory struct {
	ID     int64 `gorm:"column:id;primary_key"  json:"id"`
	BookID int64 `gorm:"column:book_id"  json:"book_id"`
	CateID int64 `gorm:"column:cate_id"  json:"cate_id"`
}
