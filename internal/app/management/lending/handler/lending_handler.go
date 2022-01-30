package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/management/pkg"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending"
	"github.com/gin-gonic/gin"
)

// LendingHandler handle the lending router call.
type LendingHandler struct {
	Module lending.Module
}

// NewLendingHandler Return an instance of this handler.
func NewLendingHandler(module lending.Module) LendingHandler {
	return LendingHandler{
		Module: module,
	}
}

// Get returns all lendings.
func (h LendingHandler) Get(c *gin.Context) {
	var params lending.Filter
	err := c.Bind(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid parameters")
	}
	if params != (lending.Filter{}) {
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

// Find return one lending by ID.
func (h LendingHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Create persist a lending to the database.
func (h LendingHandler) Create(c *gin.Context) {
	lend, apiError := pkg.AssociateLendingInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(lend)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Update update an existent lending.
func (h LendingHandler) Update(c *gin.Context) {
	upLending, apiError := pkg.AssociateLendingInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upLending)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}

// Delete delete an existent lending.
func (h LendingHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	c.Status(http.StatusNoContent)
}
