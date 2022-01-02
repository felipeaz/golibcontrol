package comments

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"time"
)

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId    uint      `json:"bookId" binding:"required"`
	UserId    uint      `json:"userId" binding:"required"`
	Text      string    `json:"text" binding:"required"`
	CreatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
}

type Module interface {
	Get(bookId string) ([]Comment, *errors.ApiError)
	Find(id string) (Comment, *errors.ApiError)
	Create(comment Comment) (*Comment, *errors.ApiError)
	Update(id string, upComment Comment) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get(bookId string) ([]Comment, *errors.ApiError)
	Find(id string) (Comment, *errors.ApiError)
	Create(comment Comment) (*Comment, *errors.ApiError)
	Update(id string, upComment Comment) *errors.ApiError
	Delete(id string) *errors.ApiError
}
