package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// CategoryRoutes initialize Category routes.
func CategoryRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, categoryHandler handler.CategoryHandler) {
	r := rg.Group("/category")

	r.GET("/", middleware.TokenAuth(), categoryHandler.Get)
	r.GET("/:id", middleware.TokenAuth(), categoryHandler.Find)
	r.POST("/", middleware.TokenAuth(), categoryHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), categoryHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), categoryHandler.Delete)
}
