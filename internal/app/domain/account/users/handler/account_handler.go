package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	Module users.Module
}

// NewAccountHandler returns an instance of authHandler
func NewAccountHandler(module users.Module) AccountHandler {
	return AccountHandler{
		Module: module,
	}
}

// Get returns all accounts.
func (h AccountHandler) Get(c *gin.Context) {
	resp, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Find return one user by ID.
func (h AccountHandler) Find(c *gin.Context) {
	resp, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Create creates a user
func (h AccountHandler) Create(c *gin.Context) {
	account, apiError := pkg.ParseAccountEntry(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	resp, apiError := h.Module.Create(account)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// Update update an existent user.
func (h AccountHandler) Update(c *gin.Context) {
	upAccount, apiError := pkg.ParseAccountEntry(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	apiError = h.Module.Update(c.Param("id"), upAccount)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete delete an existent user by id.
func (h AccountHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.Status(http.StatusNoContent)
}
