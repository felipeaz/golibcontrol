package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/pkg"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/module/interface"
	"github.com/gin-gonic/gin"
)

type ReserveHandler struct {
	Module _interface.ReserveModuleInterface
}

func NewReserveHandler(module _interface.ReserveModuleInterface) ReserveHandler {
	return ReserveHandler{
		Module: module,
	}
}

func (h ReserveHandler) Get(c *gin.Context) {
	reserves, apiError := h.Module.Get(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reserves})
}

func (h ReserveHandler) Find(c *gin.Context) {
	reserve, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reserve})
}

func (h ReserveHandler) Create(c *gin.Context) {
	reserve, apiError := pkg.AssociateReserveInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	id, apiError := h.Module.Create(reserve)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
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
