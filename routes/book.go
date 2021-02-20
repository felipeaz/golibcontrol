package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addBookRoutes(r *gin.RouterGroup) {
	book := r.Group("book")

	book.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Book Route!",
		})
	})
}
