package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments/model/converter"
)

type CommentRepository struct {
	DB database.GORMServiceInterface
}

func NewCommentRepository(db database.GORMServiceInterface) CommentRepository {
	return CommentRepository{
		DB: db,
	}
}

func (r CommentRepository) Get(bookId string) ([]model.Comment, *errors.ApiError) {
	result, err := r.DB.FetchAllWhere(&[]model.Comment{}, "book_id", bookId)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	comments, apiError := converter.ConvertToSliceCommentObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return comments, nil
}

func (r CommentRepository) Find(id string) (model.Comment, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Comment{}, id)
	if err != nil {
		return model.Comment{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	comment, apiError := converter.ConvertToCommentObj(result)
	if apiError != nil {
		return model.Comment{}, apiError
	}
	return comment, nil
}

func (r CommentRepository) Create(comment model.Comment) (*model.Comment, *errors.ApiError) {
	err := r.DB.Persist(&comment)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &comment, nil
}

func (r CommentRepository) Update(id string, upComment model.Comment) *errors.ApiError {
	err := r.DB.Refresh(&upComment, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r CommentRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&model.Comment{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
