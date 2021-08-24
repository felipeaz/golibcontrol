package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/handler"
	"github.com/gin-gonic/gin"
)

func ReserveRoutes(rg *gin.RouterGroup, reserveHandler handler.ReserveHandler) {
	r := rg.Group("/reserve")
	r.GET("/:bookId", reserveHandler.Get)
	r.POST("/", reserveHandler.Create)
	r.PUT("/:id", reserveHandler.Update)
	r.DELETE("/:id", reserveHandler.Delete)
}
