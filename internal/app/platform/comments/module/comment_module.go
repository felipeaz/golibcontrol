package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type CommentModule struct {
	Repository comments.Repository
	Log        logger.LogInterface
}

func NewCommentModule(repo comments.Repository, log logger.LogInterface) CommentModule {
	return CommentModule{
		Repository: repo,
		Log:        log,
	}
}

func (m CommentModule) Get(bookId string) ([]comments.Comment, *errors.ApiError) {
	return m.Repository.Get(bookId)
}

func (m CommentModule) Find(id string) (comments.Comment, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m CommentModule) Create(comment comments.Comment) (*comments.Comment, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m CommentModule) Update(id string, upComment comments.Comment) *errors.ApiError {
	return m.Repository.Update(id, upComment)
}

func (m CommentModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
