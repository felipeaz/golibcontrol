package handler

import (
	"net/http"

	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups/module/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/pkg"
	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	Module _interface.GroupModuleInterface
}

func NewGroupHandler(module _interface.GroupModuleInterface) GroupHandler {
	return GroupHandler{
		Module: module,
	}
}

func (h GroupHandler) Get(c *gin.Context) {
	groups, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": groups})
}

func (h GroupHandler) Find(c *gin.Context) {
	group, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": group})
}

func (h GroupHandler) Create(c *gin.Context) {
	group, apiError := pkg.AssociateGroupInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	id, apiError := h.Module.Create(group)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
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
