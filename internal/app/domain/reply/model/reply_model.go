package model

import (
	"time"
)

type Reply struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserId    uint      `json:"userId"`
	CommentId uint      `json:"commentId"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}
