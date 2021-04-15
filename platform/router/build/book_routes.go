package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// BookRoutes initialize Book routes.
func BookRoutes(rg *gin.RouterGroup, bookHandler rest.BookHandler) {
	r := rg.Group("/book")

	r.GET("/", bookHandler.Get)
	r.GET("/:id", bookHandler.Find)
	r.POST("/", bookHandler.Create)
	r.PUT("/:id", bookHandler.Update)
	r.DELETE("/:id", bookHandler.Delete)
}
