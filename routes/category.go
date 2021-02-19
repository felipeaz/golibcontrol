package routes

import "github.com/gin-gonic/gin"

func addCategoryRoute(r *gin.RouterGroup) {
	category := r.Group("category")

	category.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Category Route!",
		})
	})
}
