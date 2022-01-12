package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies"
	"github.com/FelipeAz/golibcontrol/internal/app/platform/replies/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
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
	tx := r.DB.Where(nil, fmt.Sprintf("comment_id = %s", bookId))
	result, err := r.DB.Find(tx, &[]replies.Reply{})
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &replies.Reply{})
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upReply)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &replies.Reply{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
