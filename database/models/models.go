package models

import (
	"time"

	"gorm.io/gorm"
)

// DB is an instance of gorm DB
var DB *gorm.DB

// Student contains all Student's table properties.
type Student struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	RegisterNumber string    `json:"registerNumber" binding:"required"`
	Name           string    `json:"name" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	Phone          string    `json:"phone" binding:"required"`
	Birth          time.Time `json:"birth" binding:"required"`
}

// Book contains all Book's table properties.
type Book struct {
	ID         uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title      string `json:"title" binding:"required"`
	Author     string `json:"author" binding:"required"`
	CategoryID uint   `json:"categoryId" binding:"required"`
}

// Category contains all Category's table properties.
type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name string `json:"name" binding:"required"`
}

// Lending contains all Lending's table properties.
type Lending struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookID      uint      `json:"bookId" binding:"required"`
	StudendID   uint      `json:"studentId" binding:"required"`
	LendingDate time.Time `json:"lendingDate" sql:"type:datetime"`
}
