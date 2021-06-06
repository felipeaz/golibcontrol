package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/model"
)

type CommentRepositoryInterface interface {
	GetComments(bookId string) ([]model.Comment, error)
	CreateComment(comment model.Comment) (uint, error)
	DeleteComment(commentId string)
}
