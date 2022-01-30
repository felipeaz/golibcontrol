package replies

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

type Reply struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CommentId uint      `json:"commentId" gorm:"unique"`
	UserId    uint      `json:"userId" binding:"required"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (r Reply) TableName() string {
	return "replies"
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
