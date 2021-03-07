package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/handler"
	"github.com/gin-gonic/gin"
)

// BookRoutes initialize Book routes.
func BookRoutes(r *gin.Engine) {
	rg := r.Group("/book")

	rg.GET("/", handler.GetBooks)
	rg.GET("/:id", handler.GetBook)
	rg.POST("/", handler.CreateBook)
	rg.PUT("/:id", handler.UpdateBook)
	rg.DELETE("/:id", handler.DeleteBook)
}
