package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth/pkg"
	userpkg "github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/pkg"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Module auth.Module
}

// NewAuthHandler returns an instance of authHandler
func NewAuthHandler(module auth.Module) AuthHandler {
	return AuthHandler{
		Module: module,
	}
}

// Login authenticate the user
func (h AuthHandler) Login(c *gin.Context) {
	credentials, apiError := userpkg.ParseAccountEntry(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}

	loginMsg := h.Module.Login(credentials)
	c.JSON(loginMsg.Status, loginMsg)
}

// Logout authenticate the user
func (h AuthHandler) Logout(c *gin.Context) {
	session, apiError := pkg.ParseSessionEntry(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	logoutMsg := h.Module.Logout(session)
	c.JSON(logoutMsg.Status, logoutMsg)
}