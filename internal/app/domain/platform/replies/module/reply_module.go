package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/model"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type ReplyModule struct {
	Repository _interface.ReplyRepositoryInterface
	Log        logger.LogInterface
}

func NewReplyModule(repo _interface.ReplyRepositoryInterface, log logger.LogInterface) ReplyModule {
	return ReplyModule{
		Repository: repo,
		Log:        log,
	}
}

func (m ReplyModule) Get(bookId string) ([]model.Reply, *errors.ApiError) {
	return m.Repository.Get(bookId)
}

func (m ReplyModule) Find(id string) (model.Reply, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ReplyModule) Create(comment model.Reply) (uint, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m ReplyModule) Update(id string, upReply model.Reply) *errors.ApiError {
	return m.Repository.Update(id, upReply)
}

func (m ReplyModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
