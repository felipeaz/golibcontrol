package model

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId    uint      `json:"bookId"`
	UserId    uint      `json:"userId"`
	Title     string    `json:"title" binding:"required"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}
