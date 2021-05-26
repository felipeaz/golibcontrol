package model

import (
	"time"
)

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId    uint      `json:"bookId"`
	UserId    uint      `json:"userId"`
	Rating    int       `json:"rating" binding:"required"`
	Title     string    `json:"title"`
	Review    string    `json:"text"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}
