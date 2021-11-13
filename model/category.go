package model

type Category struct {
	ID   int64  `gorm:"column:id;primary_key"  json:"id"`
	Name string `gorm:"column:name"  json:"name"`
}
