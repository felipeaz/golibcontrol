package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type CommentModule struct {
	Repository _interface.CommentRepositoryInterface
	Log        logger.LogInterface
}

func NewCommentModule(repo _interface.CommentRepositoryInterface, log logger.LogInterface) CommentModule {
	return CommentModule{
		Repository: repo,
		Log:        log,
	}
}

func (m CommentModule) Get(bookId string) ([]model.Comment, *errors.ApiError) {
	return m.Repository.Get(bookId)
}

func (m CommentModule) Find(id string) (model.Comment, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m CommentModule) Create(comment model.Comment) (uint, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m CommentModule) Update(id string, upComment model.Comment) *errors.ApiError {
	return m.Repository.Update(id, upComment)
}

func (m CommentModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
