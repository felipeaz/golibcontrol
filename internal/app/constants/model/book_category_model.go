package model

type BookCategory struct {
	ID         uint `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookID     uint `json:"book_id" gorm:"not null"`
	CategoryID uint `json:"category_id" gorm:"not null"`
}
