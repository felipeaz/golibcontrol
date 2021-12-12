package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments/handler"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(rg *gin.RouterGroup, commentHandler handler.CommentHandler) {
	r := rg.Group("/comments")
	r.GET("/:id/book", commentHandler.Get)
	r.GET("/:id", commentHandler.Find)
	r.POST("", commentHandler.Create)
	r.PUT("/:id", commentHandler.Update)
	r.DELETE("/:id", commentHandler.Delete)
}
