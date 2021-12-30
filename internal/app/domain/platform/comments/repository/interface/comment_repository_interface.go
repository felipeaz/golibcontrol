package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments/model"
)

type CommentRepositoryInterface interface {
	Get(bookId string) ([]model.Comment, *errors.ApiError)
	Find(id string) (model.Comment, *errors.ApiError)
	Create(comment model.Comment) (*model.Comment, *errors.ApiError)
	Update(id string, upComment model.Comment) *errors.ApiError
	Delete(id string) *errors.ApiError
}
