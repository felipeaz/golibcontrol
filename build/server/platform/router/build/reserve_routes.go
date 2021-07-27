package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func ReserveRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, reserveHandler handler.ReserveHandler) {
	r := rg.Group("/reserve")
	r.GET("/:bookId", middleware.TokenAuth(), reserveHandler.Get)
	r.POST("/", middleware.TokenAuth(), reserveHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), reserveHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), reserveHandler.Delete)
}
