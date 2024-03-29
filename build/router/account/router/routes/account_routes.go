package routes

import (
	"github.com/FelipeAz/golibcontrol/internal/app/account/users/handler"
	"github.com/gin-gonic/gin"
)

// UserRoutes initialize Account routes.
func UserRoutes(rg *gin.RouterGroup, accountHandler handler.AccountHandler) {
	r := rg.Group("/users")
	r.GET("", accountHandler.Get)
	r.GET("/:id", accountHandler.Find)
	r.PUT("/:id", accountHandler.Update)
	r.DELETE("/:id", accountHandler.Delete)

	r = rg.Group("/signin")
	r.POST("", accountHandler.Create)
}
