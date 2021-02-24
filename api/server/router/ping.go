package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addPingRoute(r *gin.RouterGroup) {
	ping := r.Group("ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
