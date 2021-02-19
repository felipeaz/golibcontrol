package routes

import "github.com/gin-gonic/gin"

func addLendingRoutes(r *gin.RouterGroup) {
	lending := r.Group("lending")

	lending.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Lending Route!",
		})
	})
}
