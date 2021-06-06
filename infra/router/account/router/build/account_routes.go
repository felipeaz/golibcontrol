package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/user/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// UserRoutes initialize Account routes.
func UserRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, accountHandler handler.AccountHandler) {
	r := rg.Group("/user")

	r.GET("/", accountHandler.Get)
	r.GET("/:id", accountHandler.Find)
	r.PUT("/:id", accountHandler.Update)
	r.DELETE("/:id", accountHandler.Delete)

	r = rg.Group("/signin")
	r.POST("/", accountHandler.Create)

	r = rg.Group("/login")
	r.POST("/", accountHandler.Login)

	r = rg.Group("/logout")
	r.POST("/", middleware.TokenAuth(), accountHandler.Logout)
}
