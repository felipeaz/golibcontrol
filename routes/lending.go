package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addLendingRoutes(r *gin.RouterGroup) {
	lending := r.Group("lending")

	lending.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Lending Route!",
		})
	})
}
