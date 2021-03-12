package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler"
)

// CategoryRoutes initialize Category routes.
func CategoryRoutes(r *gin.Engine, categoryHandler handler.CategoryHandler) {
	rg := r.Group("/category")

	rg.GET("/", categoryHandler.Get)
	rg.GET("/:id", categoryHandler.Find)
	rg.POST("/", categoryHandler.Create)
	rg.PUT("/:id", categoryHandler.Update)
	rg.DELETE("/:id", categoryHandler.Delete)
}
