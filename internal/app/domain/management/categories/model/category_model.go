package model

import "time"

// Category contains all Category's table properties.
type Category struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `time_format:"2006-01-02 15:04:05"`
}
