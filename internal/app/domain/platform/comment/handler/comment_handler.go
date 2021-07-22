package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/module"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/repository"
	"github.com/FelipeAz/golibcontrol/internal/pkg"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	Module module.CommentModule
}

func NewCommentHandler(dbService *service.MySQLService) CommentHandler {
	return CommentHandler{
		Module: module.NewCommentModule(repository.NewCommentRepository(dbService)),
	}
}

func (h CommentHandler) Get(c *gin.Context) {
	comments, apiError := h.Module.Get(c.Param("bookId"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comments})
}

func (h CommentHandler) Find(c *gin.Context) {
	comment, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

func (h CommentHandler) Create(c *gin.Context) {
	comment, apiError := pkg.AssociateCommentInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	id, apiError := h.Module.Create(comment)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h CommentHandler) Update(c *gin.Context) {
	upComment, apiError := pkg.AssociateCommentInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upComment)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h CommentHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}
