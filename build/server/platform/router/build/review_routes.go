package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func ReviewRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, reviewHandler handler.ReviewHandler) {
	r := rg.Group("/review")
	r.GET("/:bookId", middleware.TokenAuth(), reviewHandler.Get)
	r.POST("/", middleware.TokenAuth(), reviewHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), reviewHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), reviewHandler.Delete)
}
