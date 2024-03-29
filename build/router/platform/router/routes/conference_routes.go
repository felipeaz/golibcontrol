package routes

import (
	"github.com/FelipeAz/golibcontrol/internal/app/platform/conferences/handler"
	"github.com/gin-gonic/gin"
)

func ConferenceRoutes(rg *gin.RouterGroup, conferenceHandler handler.ConferenceHandler) {
	r := rg.Group("/conferences")
	r.GET("", conferenceHandler.Get)
	r.GET("/:id", conferenceHandler.Find)
	r.POST("", conferenceHandler.Create)
	r.POST("/:id/subscribe", conferenceHandler.Subscribe)
	r.POST("/:id/unsubscribe", conferenceHandler.Unsubscribe)
	r.PUT("/:id", conferenceHandler.Update)
	r.DELETE("/:id", conferenceHandler.Delete)
}
