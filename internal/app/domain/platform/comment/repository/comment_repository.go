package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func (r CommentRepository) GetComments(bookId string) ([]model.Comment, error) {
	return nil, nil
}

func (r CommentRepository) CreateComment(comment model.Comment) (uint, error) {
	return 0, nil
}

func (r CommentRepository) DeleteComment(commentId string) {

}
