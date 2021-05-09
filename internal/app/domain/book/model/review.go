package model

import (
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
)

type Review struct {
	ID        uint          `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId    Book          `json:"bookId" binding:"required"`
	UserId    model.Account `json:"userId" binding:"required"`
	Rating    int           `json:"rating" binding:"required"`
	Title     string        `json:"title"`
	Review    string        `json:"text"`
	CreatedAt time.Time     `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time     `time_format:"2006-01-02 15:04:05"`
}
