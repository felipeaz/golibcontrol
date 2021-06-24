package handler

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/infra/jwt"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/infra/redis"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/module"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/repository"
	"github.com/FelipeAz/golibcontrol/internal/pkg"
	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	Module module.AccountModule
}

// NewAccountHandler returns an instance of authHandler
func NewAccountHandler(auth *jwt.Auth, dbService *service.MySQLService, cache *redis.Cache) AccountHandler {
	return AccountHandler{
		Module: module.AccountModule{
			Repository: repository.AccountRepository{
				DB: dbService,
			},
			Cache: cache,
			Auth:  auth,
		},
	}
}

// Login authenticate the user
func (h AccountHandler) Login(c *gin.Context) {
	credentials, err := pkg.AssociateAccountInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	loginMsg := h.Module.Login(credentials)
	c.JSON(loginMsg.Status, loginMsg)
}

// Logout authenticate the user
func (h AccountHandler) Logout(c *gin.Context) {
	logoutMsg := h.Module.Logout(c.Request)
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

// Create creates an user
func (h AccountHandler) Create(c *gin.Context) {
	account, err := pkg.AssociateAccountInput(c)
	if err != nil {
		c.JSON(err.Status, err)
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
	upAccount, err := pkg.AssociateAccountInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	apiError := h.Module.Update(c.Param("id"), upAccount)
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
