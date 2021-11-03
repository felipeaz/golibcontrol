package model

type BookCategory struct {
	ID         uint `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookID     uint `json:"bookId" gorm:"not null"`
	CategoryID uint `json:"categoryId" gorm:"not null"`
}
