package table

import "time"

// Student contains all Student's table properties.
type Student struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	RegisterNumber string    `json:"registerNumber"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	Birth          time.Time `json:"birth"`
}

// Book contains all Book's table properties.
type Book struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	CategoryID uint   `json:"categoryId"`
}

// Category contains all Category's table properties.
type Category struct {
	ID    uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title string `json:"title"`
}

// Lending contains all Lending's table properties.
type Lending struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookID      uint      `json:"bookId"`
	StudendID   uint      `json:"studentId"`
	LendingDate time.Time `json:"lendingDate" sql:"type:datetime"`
}
