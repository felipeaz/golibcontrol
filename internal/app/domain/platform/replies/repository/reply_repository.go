package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/pkg"
)

type ReplyRepository struct {
	DB database.GORMServiceInterface
}

func NewReplyRepository(db database.GORMServiceInterface) ReplyRepository {
	return ReplyRepository{
		DB: db,
	}
}

func (r ReplyRepository) Get(bookId string) ([]replies.Reply, *errors.ApiError) {
	result, err := r.DB.FetchAllWhere(&[]replies.Reply{}, "comment_id", bookId)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParserToSliceReplyObj(result)
}

func (r ReplyRepository) Find(id string) (replies.Reply, *errors.ApiError) {
	result, err := r.DB.Fetch(&replies.Reply{}, id)
	if err != nil {
		return replies.Reply{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParserToReplyObj(result)
}

func (r ReplyRepository) Create(reply replies.Reply) (*replies.Reply, *errors.ApiError) {
	err := r.DB.Persist(&reply)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &reply, nil
}

func (r ReplyRepository) Update(id string, upReply replies.Reply) *errors.ApiError {
	err := r.DB.Refresh(&upReply, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r ReplyRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&replies.Reply{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
