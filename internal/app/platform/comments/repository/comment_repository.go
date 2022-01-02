package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments"
	"github.com/FelipeAz/golibcontrol/internal/app/platform/comments/pkg"
)

type CommentRepository struct {
	DB database.GORMServiceInterface
}

func NewCommentRepository(db database.GORMServiceInterface) CommentRepository {
	return CommentRepository{
		DB: db,
	}
}

func (r CommentRepository) Get(bookId string) ([]comments.Comment, *errors.ApiError) {
	result, err := r.DB.FetchAllWhere(&[]comments.Comment{}, "book_id", bookId)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceCommentObj(result)
}

func (r CommentRepository) Find(id string) (comments.Comment, *errors.ApiError) {
	result, err := r.DB.Fetch(&comments.Comment{}, id)
	if err != nil {
		return comments.Comment{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToCommentObj(result)
}

func (r CommentRepository) Create(comment comments.Comment) (*comments.Comment, *errors.ApiError) {
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

func (r CommentRepository) Update(id string, upComment comments.Comment) *errors.ApiError {
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
	err := r.DB.Remove(&comments.Comment{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
