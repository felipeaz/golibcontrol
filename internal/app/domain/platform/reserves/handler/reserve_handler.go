package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves"
	"github.com/gin-gonic/gin"
)

type ReserveHandler struct {
	Module reserves.Module
}

func NewReserveHandler(module reserves.Module) ReserveHandler {
	return ReserveHandler{
		Module: module,
	}
}

func (h ReserveHandler) Get(c *gin.Context) {
	resp, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h ReserveHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h ReserveHandler) Create(c *gin.Context) {
	reserve, apiError := pkg.AssociateReserveInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(reserve)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h ReserveHandler) Update(c *gin.Context) {
	upReserve, apiError := pkg.AssociateReserveInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upReserve)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}

func (h ReserveHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}
