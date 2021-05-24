package model

import (
	"time"

	accountModel "github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	bookModel "github.com/FelipeAz/golibcontrol/internal/app/domain/book/model"
)

type Reserve struct {
	ID         uint                 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId     bookModel.Book       `json:"bookId" binding:"required"`
	UserId     accountModel.Account `json:"userId" binding:"required"`
	CreatedAt  time.Time            `time_format:"2006-01-02 15:04:05"`
	RetrieveAt time.Time            `time_format:"2006-01-02 15:04:05"`
}
