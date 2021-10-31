package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/handler"
	"github.com/gin-gonic/gin"
)

func ReviewRoutes(rg *gin.RouterGroup, reviewHandler handler.ReviewHandler) {
	r := rg.Group("/reviews")
	r.GET("/:bookId", reviewHandler.Get)
	r.POST("/", reviewHandler.Create)
	r.PUT("/:id", reviewHandler.Update)
	r.DELETE("/:id", reviewHandler.Delete)
}
