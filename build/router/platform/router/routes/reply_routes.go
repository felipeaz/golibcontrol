package routes

import (
	"github.com/FelipeAz/golibcontrol/internal/app/platform/replies/handler"
	"github.com/gin-gonic/gin"
)

func ReplyRoutes(rg *gin.RouterGroup, replyHandler handler.ReplyHandler) {
	r := rg.Group("/replies")
	r.GET("/:id/comment", replyHandler.Get)
	r.GET("/:id", replyHandler.Find)
	r.POST("/:id/comment", replyHandler.Create)
	r.PUT("/:id", replyHandler.Update)
	r.DELETE("/:id", replyHandler.Delete)
}
