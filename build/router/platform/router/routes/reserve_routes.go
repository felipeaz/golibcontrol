package routes

import (
	"github.com/FelipeAz/golibcontrol/internal/app/platform/reserves/handler"
	"github.com/gin-gonic/gin"
)

func ReserveRoutes(rg *gin.RouterGroup, reserveHandler handler.ReserveHandler) {
	r := rg.Group("/reserves")
	r.GET("", reserveHandler.Get)
	r.GET("/:id", reserveHandler.Find)
	r.POST("", reserveHandler.Create)
	r.PUT("/:id", reserveHandler.Update)
	r.DELETE("/:id", reserveHandler.Delete)
}
