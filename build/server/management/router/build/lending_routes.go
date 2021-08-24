package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/handler"
	"github.com/gin-gonic/gin"
)

// LendingRoutes initialize Category routes.
func LendingRoutes(rg *gin.RouterGroup, lendingHandler handler.LendingHandler) {
	r := rg.Group("/lending")
	r.GET("/", lendingHandler.Get)
	r.GET("/:id", lendingHandler.Find)
	r.POST("/", lendingHandler.Create)
	r.PUT("/:id", lendingHandler.Update)
	r.DELETE("/:id", lendingHandler.Delete)
}
