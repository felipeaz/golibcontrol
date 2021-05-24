package repository

import (
	commentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/comment/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func (r CommentRepository) GetComments(bookId string) ([]commentModel.Comment, error) {
	return nil, nil
}

func (r CommentRepository) CreateComment(comment commentModel.Comment) (uint, error) {
	return 0, nil
}

func (r CommentRepository) DeleteComment(commentId string) {

}
