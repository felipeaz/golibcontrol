package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments"
	"github.com/FelipeAz/golibcontrol/internal/app/platform/comments/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
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
	tx := r.DB.Preload("Reply")
	tx = r.DB.Where(tx, fmt.Sprintf("book_id = %s", bookId))
	result, err := r.DB.Find(tx, &[]comments.Comment{})
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
	tx := r.DB.Preload("Reply")
	tx = r.DB.Where(tx, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &comments.Comment{})
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upComment)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &comments.Comment{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
