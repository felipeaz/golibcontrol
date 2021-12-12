package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/categories/handler"
	"github.com/gin-gonic/gin"
)

// CategoryRoutes initialize Category routes.
func CategoryRoutes(rg *gin.RouterGroup, categoryHandler handler.CategoryHandler) {
	r := rg.Group("/categories")
	r.GET("", categoryHandler.Get)
	r.GET("/:id", categoryHandler.Find)
	r.POST("", categoryHandler.Create)
	r.PUT("/:id", categoryHandler.Update)
	r.DELETE("/:id", categoryHandler.Delete)
}
