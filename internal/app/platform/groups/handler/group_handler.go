package handler

import (
	domainpkg "github.com/FelipeAz/golibcontrol/internal/app/domain/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/platform/pkg"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups"
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
	var params groups.Filter
	err := c.Bind(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid parameters")
	}
	if params != (groups.Filter{}) {
		resp, apiError := h.Module.GetByFilter(params)
		if apiError != nil {
			c.JSON(apiError.Status, apiError)
			return
		}

		c.JSON(http.StatusOK, resp)
		return
	}
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

func (h GroupHandler) Subscribe(c *gin.Context) {
	subscription, apiError := pkg.AssociateGroupSubscriberInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	id, err := domainpkg.ParseStringToId(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	subscription.GroupID = id
	resp, apiError := h.Module.Subscribe(subscription)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h GroupHandler) Unsubscribe(c *gin.Context) {
	subscription, apiError := pkg.AssociateGroupSubscriberInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	id, err := domainpkg.ParseStringToId(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	subscription.GroupID = id
	apiError = h.Module.Unsubscribe(subscription)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
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
