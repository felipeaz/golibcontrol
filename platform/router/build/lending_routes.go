package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// LendingRoutes initialize Category routes.
func LendingRoutes(r *gin.Engine, lendingHandler rest.LendingHandler) {
	rg := r.Group("/lending")

	rg.GET("/", lendingHandler.Get)
	rg.GET("/:id", lendingHandler.Find)
	rg.POST("/", lendingHandler.Create)
	rg.PUT("/:id", lendingHandler.Update)
	rg.DELETE("/:id", lendingHandler.Delete)
}
