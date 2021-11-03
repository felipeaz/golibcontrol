package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/model"
)

type ReplyRepositoryInterface interface {
	Get(bookId string) ([]model.Reply, *errors.ApiError)
	Find(id string) (model.Reply, *errors.ApiError)
	Create(reply model.Reply) (uint, *errors.ApiError)
	Update(id string, upReply model.Reply) *errors.ApiError
	Delete(id string) *errors.ApiError
}