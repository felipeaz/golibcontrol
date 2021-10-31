package handler

import (
	"net/http"

	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/account/module/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/pkg"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	Module _interface.AccountModuleInterface
}

// NewAccountHandler returns an instance of authHandler
func NewAccountHandler(module _interface.AccountModuleInterface) AccountHandler {
	return AccountHandler{
		Module: module,
	}
}

// Login authenticate the user
func (h AccountHandler) Login(c *gin.Context) {
	credentials, apiError := pkg.AssociateAccountInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	loginMsg := h.Module.Login(credentials)
	c.JSON(loginMsg.Status, loginMsg)
}

// Logout authenticate the user
func (h AccountHandler) Logout(c *gin.Context) {
	session, apiError := pkg.AssociateSessionInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	logoutMsg := h.Module.Logout(session)
	c.JSON(logoutMsg.Status, logoutMsg)
}

// Get returns all accounts.
func (h AccountHandler) Get(c *gin.Context) {
	accounts, apiError := h.Module.Get()
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": accounts})
}

// Find return one user by ID.
func (h AccountHandler) Find(c *gin.Context) {
	account, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}

// Create creates a user
func (h AccountHandler) Create(c *gin.Context) {
	account, apiError := pkg.AssociateAccountInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	id, apiError := h.Module.Create(account)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// Update update an existent user.
func (h AccountHandler) Update(c *gin.Context) {
	upAccount, apiError := pkg.AssociateAccountInput(c)
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
