package replies

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"time"
)

type Reply struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	CommentId uint      `json:"commentId" gorm:"unique"`
	UserId    uint      `json:"userId" binding:"required"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

type Module interface {
	Get(bookId string) ([]Reply, *errors.ApiError)
	Find(id string) (Reply, *errors.ApiError)
	Create(comment Reply) (*Reply, *errors.ApiError)
	Update(id string, upReply Reply) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get(bookId string) ([]Reply, *errors.ApiError)
	Find(id string) (Reply, *errors.ApiError)
	Create(reply Reply) (*Reply, *errors.ApiError)
	Update(id string, upReply Reply) *errors.ApiError
	Delete(id string) *errors.ApiError
}
