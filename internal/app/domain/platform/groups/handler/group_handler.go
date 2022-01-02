package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/pkg"
	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	Module groups.Module
}

func NewGroupHandler(module groups.Module) GroupHandler {
	return GroupHandler{
		Module: module,
	}
}

func (h GroupHandler) Get(c *gin.Context) {
	resp, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h GroupHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h GroupHandler) Create(c *gin.Context) {
	group, apiError := pkg.AssociateGroupInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(group)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h GroupHandler) Update(c *gin.Context) {
	upGroup, apiError := pkg.AssociateGroupInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upGroup)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h GroupHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}
