package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/platform/pkg"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	Module comments.Module
}

func NewCommentHandler(module comments.Module) CommentHandler {
	return CommentHandler{
		Module: module,
	}
}

func (h CommentHandler) Get(c *gin.Context) {
	resp, apiError := h.Module.Get(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h CommentHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h CommentHandler) Create(c *gin.Context) {
	comment, apiError := pkg.AssociateCommentInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(comment)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
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
