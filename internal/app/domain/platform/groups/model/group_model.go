package model

import (
	"time"
)

type Group struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Status      string    `json:"status" binding:"required"`
	CreatedAt   time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `time_format:"2006-01-02 15:04:05"`
}
