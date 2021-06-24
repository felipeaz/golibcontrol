package handler

import (
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/module"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/repository"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	Module module.CommentModule
}

func NewCommentHandler(dbService *service.MySQLService) CommentHandler {
	return CommentHandler{
		Module: module.CommentModule{
			Repository: repository.CommentRepository{
				DB: dbService,
			},
		},
	}
}

func (h CommentHandler) Get(c *gin.Context) {

}

func (h CommentHandler) Find(c *gin.Context) {

}

func (h CommentHandler) Create(c *gin.Context) {

}

func (h CommentHandler) Update(c *gin.Context) {

}

func (h CommentHandler) Delete(c *gin.Context) {

}
