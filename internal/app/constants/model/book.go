package model

import "time"

// Book contains all Book's table properties.
type Book struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Title          string    `json:"title" binding:"required"`
	Author         string    `json:"author" binding:"required"`
	RegisterNumber string    `json:"registerNumber" binding:"required" gorm:"unique"`
	Available      bool      `json:"available" gorm:"default:true"`
	CreatedAt      time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt      time.Time `time_format:"2006-01-02 15:04:05"`
}
