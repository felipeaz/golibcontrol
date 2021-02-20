package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addStudentRoutes(r *gin.RouterGroup) {
	student := r.Group("student")

	student.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Student Route!",
		})
	})
}
