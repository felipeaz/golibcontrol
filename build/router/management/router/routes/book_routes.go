package routes

import (
	"github.com/FelipeAz/golibcontrol/internal/app/management/books/handler"
	"github.com/gin-gonic/gin"
)

// BookRoutes initialize Book routes.
func BookRoutes(rg *gin.RouterGroup, bookHandler handler.BookHandler) {
	r := rg.Group("/book")
	r.GET("", bookHandler.Get)
	r.GET("/:id", bookHandler.Find)
	r.POST("", bookHandler.Create)
	r.PUT("/:id", bookHandler.Update)
	r.DELETE("/:id", bookHandler.Delete)
}
