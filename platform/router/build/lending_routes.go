package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// LendingRoutes initialize Category routes.
func LendingRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, lendingHandler rest.LendingHandler) {
	r := rg.Group("/lending")

	r.GET("/", middleware.TokenAuth(), lendingHandler.Get)
	r.GET("/:id", middleware.TokenAuth(), lendingHandler.Find)
	r.POST("/", middleware.TokenAuth(), lendingHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), lendingHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), lendingHandler.Delete)
}
