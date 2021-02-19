package routes

import "github.com/gin-gonic/gin"

func addStudentRoutes(r *gin.RouterGroup) {
	student := r.Group("student")

	student.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Student Route!",
		})
	})
}
