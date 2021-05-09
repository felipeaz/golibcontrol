package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/book/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// BookRoutes initialize Book routes.
func BookRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, bookHandler handler.BookHandler) {
	r := rg.Group("/book")

	r.GET("/", middleware.TokenAuth(), bookHandler.Get)
	r.GET("/:id", middleware.TokenAuth(), bookHandler.Find)
	r.POST("/", middleware.TokenAuth(), bookHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), bookHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), bookHandler.Delete)
}
