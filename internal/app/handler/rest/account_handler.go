package rest

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/module"
	"github.com/FelipeAz/golibcontrol/internal/app/repository"
	"github.com/FelipeAz/golibcontrol/internal/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountHandler struct {
	Module module.AccountModule
}

// NewAuthHandler returns an instance of authHandler
func NewAccountHandler(db *gorm.DB) AccountHandler {
	return AccountHandler{
		Module: module.AccountModule{
			Repository: repository.AccountRepository{
				DB: db,
			},
		},
	}
}

// Login authenticate the user
func (h AccountHandler) Login(c *gin.Context) {
	credentials, err := pkg.AssociateLoginInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	loginMsg := h.Module.Login(credentials)
	c.JSON(loginMsg.Status, loginMsg)
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

// Find return one account by ID.
func (h AccountHandler) Find(c *gin.Context) {
	account, apiError := h.Module.Find(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}

// Create creates an account
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

// Update update an existent account.
func (h AccountHandler) Update(c *gin.Context) {
	upAccount, err := pkg.AssociateAccountInput(c)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	account, apiError := h.Module.Update(c.Param("id"), upAccount)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": account})
}

// Delete delete an existent account by id.
func (h AccountHandler) Delete(c *gin.Context) {
	apiError := h.Module.Delete(c.Param("id"))
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}
