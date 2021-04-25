package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// CategoryRoutes initialize Category routes.
func CategoryRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, categoryHandler rest.CategoryHandler) {
	r := rg.Group("/category")

	r.GET("/", middleware.TokenAuth(), categoryHandler.Get)
	r.GET("/:id", middleware.TokenAuth(), categoryHandler.Find)
	r.POST("/", middleware.TokenAuth(), categoryHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), categoryHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), categoryHandler.Delete)
}
