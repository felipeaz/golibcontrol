package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/handler"
	"github.com/gin-gonic/gin"
)

func CommentsRoutes(rg *gin.RouterGroup, commentHandler handler.CommentHandler) {
	r := rg.Group("/comments")
	r.GET("/:bookId", commentHandler.Get)
	r.POST("/", commentHandler.Create)
	r.PUT("/:id", commentHandler.Update)
	r.DELETE("/:id", commentHandler.Delete)
}
