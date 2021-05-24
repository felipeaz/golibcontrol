package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/comment/module"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/comment/repository"
	"gorm.io/gorm"
)

type CommentHandler struct {
	Module module.CommentModule
}

func NewCommentHandler(db *gorm.DB) CommentHandler {
	return CommentHandler{
		Module: module.CommentModule{
			Repository: repository.CommentRepository{
				DB: db,
			},
		},
	}
}
