package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/model"
)

type ReplyModuleInterface interface {
	Get(bookId string) ([]model.Reply, *errors.ApiError)
	Find(id string) (model.Reply, *errors.ApiError)
	Create(comment model.Reply) (*model.Reply, *errors.ApiError)
	Update(id string, upReply model.Reply) *errors.ApiError
	Delete(id string) *errors.ApiError
}
