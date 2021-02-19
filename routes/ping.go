package routes

import "github.com/gin-gonic/gin"

func addPingRoute(r *gin.RouterGroup) {
	ping := r.Group("ping")

	ping.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
