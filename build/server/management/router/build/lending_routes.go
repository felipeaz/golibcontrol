package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// LendingRoutes initialize Category routes.
func LendingRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, lendingHandler handler.LendingHandler) {
	r := rg.Group("/lending")
	r.GET("/", middleware.TokenAuth(), lendingHandler.Get)
	r.GET("/:id", middleware.TokenAuth(), lendingHandler.Find)
	r.POST("/", middleware.TokenAuth(), lendingHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), lendingHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), lendingHandler.Delete)
}
