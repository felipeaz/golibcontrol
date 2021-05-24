package _interface

import (
	commentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/comment/model"
)

type CommentRepositoryInterface interface {
	GetComments(bookId string) ([]commentModel.Comment, error)
	CreateComment(comment commentModel.Comment) (uint, error)
	DeleteComment(commentId string)
}
