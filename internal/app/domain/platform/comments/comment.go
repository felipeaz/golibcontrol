package comments

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	BookId uint   `json:"bookId" binding:"required"`
	UserId uint   `json:"userId" binding:"required"`
	Text   string `json:"text" binding:"required"`
}

func (c Comment) TableName() string {
	return "comments"
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
