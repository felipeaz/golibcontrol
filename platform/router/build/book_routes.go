package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// BookRoutes initialize Book routes.
func BookRoutes(r *gin.Engine, bookHandler rest.BookHandler) {
	rg := r.Group("/book")

	rg.GET("/", bookHandler.Get)
	rg.GET("/:id", bookHandler.Find)
	rg.POST("/", bookHandler.Create)
	rg.PUT("/:id", bookHandler.Update)
	rg.DELETE("/:id", bookHandler.Delete)
}
