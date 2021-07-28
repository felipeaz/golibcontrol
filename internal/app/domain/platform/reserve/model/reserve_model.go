package model

import (
	"time"
)

type Reserve struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId     uint      `json:"bookId" gorm:"unique"`
	UserId     uint      `json:"userId" gorm:"unique"`
	RetrieveAt time.Time `json:"retrieveAt" time_format:"2006-01-02 15:04:05"`
	CreatedAt  time.Time `time_format:"2006-01-02 15:04:05"`
}
