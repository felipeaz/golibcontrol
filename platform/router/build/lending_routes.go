package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// LendingRoutes initialize Category routes.
func LendingRoutes(rg *gin.RouterGroup, lendingHandler rest.LendingHandler) {
	r := rg.Group("/lending")

	r.GET("/", lendingHandler.Get)
	r.GET("/:id", lendingHandler.Find)
	r.POST("/", lendingHandler.Create)
	r.PUT("/:id", lendingHandler.Update)
	r.DELETE("/:id", lendingHandler.Delete)
}
