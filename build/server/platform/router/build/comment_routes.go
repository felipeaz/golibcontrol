package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, commentHandler handler.CommentHandler) {
	r := rg.Group("/comment")
	r.GET("/", middleware.TokenAuth(), commentHandler.Get)
	r.GET("/:id", middleware.TokenAuth(), commentHandler.Find)
	r.POST("/", middleware.TokenAuth(), commentHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), commentHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), commentHandler.Delete)
}
