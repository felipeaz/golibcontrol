package routes

import (
	"github.com/FelipeAz/golibcontrol/internal/app/platform/groups/handler"
	"github.com/gin-gonic/gin"
)

func GroupRoutes(rg *gin.RouterGroup, groupHandler handler.GroupHandler) {
	r := rg.Group("/groups")
	r.GET("", groupHandler.Get)
	r.GET("/:id", groupHandler.Find)
	r.POST("", groupHandler.Create)
	r.PUT("/:id", groupHandler.Update)
	r.DELETE("/:id", groupHandler.Delete)
}