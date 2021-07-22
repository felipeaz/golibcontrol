package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/model/converter"
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
	result, apiError := r.DB.FindWhere(&[]model.Comment{}, "book_id", bookId)
	if apiError != nil {
		return nil, apiError
	}
	comments, apiError := converter.ConvertToSliceCommentObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return comments, nil
}

func (r CommentRepository) Find(id string) (model.Comment, *errors.ApiError) {
	result, apiError := r.DB.Find(&model.Comment{}, id)
	if apiError != nil {
		return model.Comment{}, apiError
	}
	comment, apiError := converter.ConvertToCommentObj(result)
	if apiError != nil {
		return model.Comment{}, apiError
	}
	return comment, nil
}

func (r CommentRepository) Create(comment model.Comment) (uint, *errors.ApiError) {
	apiError := r.DB.Create(&comment)
	if apiError != nil {
		return 0, apiError
	}
	return comment.ID, nil
}

func (r CommentRepository) Update(id string, upComment model.Comment) *errors.ApiError {
	return r.DB.Update(&upComment, id)
}

func (r CommentRepository) Delete(id string) *errors.ApiError {
	return r.DB.Delete(&model.Comment{}, id)
}
