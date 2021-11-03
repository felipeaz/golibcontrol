package handler

import (
	"net/http"

	_pkg "github.com/FelipeAz/golibcontrol/internal/app/domain/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/pkg"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/module/interface"
	"github.com/gin-gonic/gin"
)

type ReplyHandler struct {
	Module _interface.ReplyModuleInterface
}

func NewReplyHandler(module _interface.ReplyModuleInterface) ReplyHandler {
	return ReplyHandler{
		Module: module,
	}
}

func (h ReplyHandler) Get(c *gin.Context) {
	replys, apiError := h.Module.Get(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": replys})
}

func (h ReplyHandler) Find(c *gin.Context) {
	reply, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reply})
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

	id, apiError := h.Module.Create(reply)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
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
