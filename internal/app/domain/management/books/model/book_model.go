package model

import (
	"time"
)

// Book contains all Book's table properties.
type Book struct {
	ID             uint           `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title          string         `json:"title" binding:"required"`
	Author         string         `json:"author" binding:"required"`
	Description    string         `json:"description"`
	Image          string         `json:"image"`
	RegisterNumber string         `json:"registerNumber" binding:"required" gorm:"unique"`
	Available      bool           `json:"available" gorm:"default:true"`
	CategoriesId   string         `json:"categoriesId,omitempty" gorm:"->"`                                       // Read Only
	BookCategory   []BookCategory `gorm:"one2many:bookCategory,->;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"` // Read Only
	CreatedAt      time.Time      `time_format:"2006-01-02 15:04:05"`
	UpdatedAt      time.Time      `time_format:"2006-01-02 15:04:05"`
}

type QueryBook struct {
	Categories *string `form:"categoryId"`
	Available  *bool   `form:"available"`
}
