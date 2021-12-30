package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth/module/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/pkg"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Module _interface.AuthModule
}

// NewAuthHandler returns an instance of authHandler
func NewAuthHandler(module _interface.AuthModule) AuthHandler {
	return AuthHandler{
		Module: module,
	}
}

// Login authenticate the user
func (h AuthHandler) Login(c *gin.Context) {
	credentials, apiError := pkg.AssociateAccountInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	loginMsg := h.Module.Login(credentials)
	c.JSON(loginMsg.Status, loginMsg)
}

// Logout authenticate the user
func (h AuthHandler) Logout(c *gin.Context) {
	session, apiError := pkg.AssociateSessionInput(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	logoutMsg := h.Module.Logout(session)
	c.JSON(logoutMsg.Status, logoutMsg)
}
