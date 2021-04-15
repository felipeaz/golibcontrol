package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// CategoryRoutes initialize Category routes.
func CategoryRoutes(rg *gin.RouterGroup, categoryHandler rest.CategoryHandler) {
	r := rg.Group("/category")

	r.GET("/", categoryHandler.Get)
	r.GET("/:id", categoryHandler.Find)
	r.POST("/", categoryHandler.Create)
	r.PUT("/:id", categoryHandler.Update)
	r.DELETE("/:id", categoryHandler.Delete)
}
