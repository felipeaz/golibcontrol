package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/handler"
	"github.com/gin-gonic/gin"
)

// BookRoutes initialize Book routes.
func BookRoutes(r *gin.Engine, bookHandler handler.BookHandler) {
	rg := r.Group("/book")

	rg.GET("/", bookHandler.Get)
	rg.GET("/:id", bookHandler.Find)
	rg.POST("/", bookHandler.Create)
	rg.PUT("/:id", bookHandler.Update)
	rg.DELETE("/:id", bookHandler.Delete)
}
