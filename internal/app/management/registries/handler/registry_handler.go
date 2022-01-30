package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/app/management/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegistryHandler handle the registry router call.
type RegistryHandler struct {
	Module registries.Module
}

// NewRegistryHandler returns an instance of registry handler.
func NewRegistryHandler(module registries.Module) RegistryHandler {
	return RegistryHandler{
		Module: module,
	}
}

// Get returns all registries.
func (h RegistryHandler) Get(c *gin.Context) {
	var params registries.Filter
	err := c.Bind(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid parameters")
	}
	if params != (registries.Filter{}) {
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

// Find return one registry by ID.
func (h RegistryHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Create persist a registry to the database.
func (h RegistryHandler) Create(c *gin.Context) {
	registry, apiError := pkg.AssociateRegistryInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(registry)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Update update an existent registry.
func (h RegistryHandler) Update(c *gin.Context) {
	upRegistry, apiError := pkg.AssociateRegistryInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upRegistry)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}

// Delete delete an existent registry.
func (h RegistryHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}
