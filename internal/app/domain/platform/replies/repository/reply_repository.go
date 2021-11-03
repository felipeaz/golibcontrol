package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/model/converter"
)

type ReplyRepository struct {
	DB database.GORMServiceInterface
}

func NewReplyRepository(db database.GORMServiceInterface) ReplyRepository {
	return ReplyRepository{
		DB: db,
	}
}

func (r ReplyRepository) Get(bookId string) ([]model.Reply, *errors.ApiError) {
	result, err := r.DB.FetchAllWhere(&[]model.Reply{}, "comment_id", bookId)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	reply, apiError := converter.ConvertToSliceReplyObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return reply, nil
}

func (r ReplyRepository) Find(id string) (model.Reply, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Reply{}, id)
	if err != nil {
		return model.Reply{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	reply, apiError := converter.ConvertToReplyObj(result)
	if apiError != nil {
		return model.Reply{}, apiError
	}
	return reply, nil
}

func (r ReplyRepository) Create(reply model.Reply) (uint, *errors.ApiError) {
	err := r.DB.Persist(&reply)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return reply.ID, nil
}

func (r ReplyRepository) Update(id string, upReply model.Reply) *errors.ApiError {
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
	err := r.DB.Remove(&model.Reply{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
