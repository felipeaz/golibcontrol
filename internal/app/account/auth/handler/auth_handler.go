package handler

import (
	"github.com/FelipeAz/golibcontrol/internal/app/account/auth/pkg"
	userPkg "github.com/FelipeAz/golibcontrol/internal/app/account/users/pkg"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth"
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
	credentials, apiError := userPkg.ParseAccountEntry(c)
	if apiError != nil {
		c.JSON(apiError.Status, apiError)
		return
	}
	authData := h.Module.Login(credentials)
	c.Header("token", authData.Token)
	c.JSON(authData.Status, gin.H{"data": authData, "token": authData.Token})
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
