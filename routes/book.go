package routes

import "github.com/gin-gonic/gin"

func addBookRoutes(r *gin.RouterGroup) {
	book := r.Group("book")

	book.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Book Route!",
		})
	})
}
