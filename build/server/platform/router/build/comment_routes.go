package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/handler"
	"github.com/gin-gonic/gin"
)

func CommentRoutes(rg *gin.RouterGroup, commentHandler handler.CommentHandler) {
	r := rg.Group("/comment")
	r.GET("/:bookId", commentHandler.Get)
	r.POST("/", commentHandler.Create)
	r.PUT("/:id", commentHandler.Update)
	r.DELETE("/:id", commentHandler.Delete)
}
