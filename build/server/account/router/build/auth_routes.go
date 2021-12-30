package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth/handler"
	"github.com/gin-gonic/gin"
)

// AuthRoutes initialize Consumer routes.
func AuthRoutes(rg *gin.RouterGroup, authHandler handler.AuthHandler) {
	r := rg.Group("/login")
	r.POST("", authHandler.Login)
	r = rg.Group("/logout")
	r.POST("", authHandler.Logout)
}
