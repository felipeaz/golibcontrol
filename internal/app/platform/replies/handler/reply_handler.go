package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/platform/pkg"
	"net/http"

	_pkg "github.com/FelipeAz/golibcontrol/internal/app/domain/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies"
	"github.com/gin-gonic/gin"
)

type ReplyHandler struct {
	Module replies.Module
}

func NewReplyHandler(module replies.Module) ReplyHandler {
	return ReplyHandler{
		Module: module,
	}
}

func (h ReplyHandler) Get(c *gin.Context) {
	resp, apiError := h.Module.Get(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h ReplyHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h ReplyHandler) Create(c *gin.Context) {
	reply, apiError := pkg.AssociateReplyInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	commentId, apiError := _pkg.ParseStringToId(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	reply.CommentId = commentId

	resp, apiError := h.Module.Create(reply)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h ReplyHandler) Update(c *gin.Context) {
	upReply, apiError := pkg.AssociateReplyInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upReply)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h ReplyHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}
