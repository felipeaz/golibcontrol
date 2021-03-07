package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/golibcontrol/handler"
	"github.com/gin-gonic/gin"
)

// BookRoutes initialize Book routes.
func BookRoutes(r *gin.Engine, bookHandler handler.BookHandler) {
	rg := r.Group("/book")

	rg.GET("/", bookHandler.GetBooks)
	rg.GET("/:id", bookHandler.GetBook)
	rg.POST("/", bookHandler.CreateBook)
	rg.PUT("/:id", bookHandler.UpdateBook)
	rg.DELETE("/:id", bookHandler.DeleteBook)
}
