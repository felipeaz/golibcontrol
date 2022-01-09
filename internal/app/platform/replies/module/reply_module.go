package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/logger"
)

type ReplyModule struct {
	Repository replies.Repository
	Log        logger.LogInterface
}

func NewReplyModule(repo replies.Repository, log logger.LogInterface) ReplyModule {
	return ReplyModule{
		Repository: repo,
		Log:        log,
	}
}

func (m ReplyModule) Get(bookId string) ([]replies.Reply, *errors.ApiError) {
	return m.Repository.Get(bookId)
}

func (m ReplyModule) Find(id string) (replies.Reply, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ReplyModule) Create(comment replies.Reply) (*replies.Reply, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m ReplyModule) Update(id string, upReply replies.Reply) *errors.ApiError {
	return m.Repository.Update(id, upReply)
}

func (m ReplyModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
