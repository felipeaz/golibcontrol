package model

import (
	"time"

	modelAccount "github.com/FelipeAz/golibcontrol/internal/app/domain/account/model"
	modelComment "github.com/FelipeAz/golibcontrol/internal/app/domain/comment/model"
)

type Reply struct {
	ID        uint                 `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	UserId    modelAccount.Account `json:"userId" binding:"required"`
	CommentId modelComment.Comment `json:"commentId" binding:"required"`
	Text      string               `json:"text" binding:"required"`
	CreatedAt time.Time            `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time            `time_format:"2006-01-02 15:04:05"`
}
