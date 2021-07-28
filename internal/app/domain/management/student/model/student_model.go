package model

import "time"

// Student contains all Student's table properties.
type Student struct {
	ID        string    `json:"id" binding:"required" gorm:"primaryKey;unique;not null"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	Grade     string    `json:"grade" binding:"required"`
	Birthday  string    `json:"birthday" binding:"required"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}
