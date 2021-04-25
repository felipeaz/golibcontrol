package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// BookRoutes initialize Book routes.
func BookRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, bookHandler rest.BookHandler) {
	r := rg.Group("/book")

	r.GET("/", middleware.TokenAuth(), bookHandler.Get)
	r.GET("/:id", middleware.TokenAuth(), bookHandler.Find)
	r.POST("/", middleware.TokenAuth(), bookHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), bookHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), bookHandler.Delete)
}
