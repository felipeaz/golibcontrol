package routes

import (
	"github.com/FelipeAz/golibcontrol/controllers"
	"github.com/gin-gonic/gin"
)

func addStudentRoutes(r *gin.RouterGroup) {
	student := r.Group("student")

	student.GET("/", controllers.GetStudents)
	student.GET("/:id", controllers.GetStudent)
	student.POST("/", controllers.CreateStudent)
	student.PUT("/:id", controllers.UpdateStudent)
	student.DELETE("/:id", controllers.DeleteStudent)
}
