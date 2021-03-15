package model

import "time"

// Student contains all Student's table properties.
type Student struct {
	ID             uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	RegisterNumber string    `json:"registerNumber" binding:"required" gorm:"unique"`
	Name           string    `json:"name" binding:"required"`
	Email          string    `json:"email" binding:"required"`
	Phone          string    `json:"phone" binding:"required"`
	Grade          string    `json:"grade" binding:"required"`
	Birthday       string    `json:"birthday" binding:"required"`
	CreatedAt      time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt      time.Time `time_format:"2006-01-02 15:04:05"`
}
