package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// AccountRoutes initialize Account routes.
func AccountRoutes(rg *gin.RouterGroup, accountHandler rest.AccountHandler) {
	r := rg.Group("/account")

	r.GET("/", accountHandler.Get)
	r.GET("/:id", accountHandler.Find)
	r.POST("/", accountHandler.Create)
	r.PUT("/:id", accountHandler.Update)
	r.DELETE("/:id", accountHandler.Delete)

	r = rg.Group("/login")
	r.GET("/", accountHandler.Login)
}
