package routes

import (
	"github.com/FelipeAz/golibcontrol/controllers"
	"github.com/gin-gonic/gin"
)

func addLendingRoutes(r *gin.RouterGroup) {
	lending := r.Group("lending")

	lending.GET("/", controllers.GetLendings)
	lending.GET("/:id", controllers.GetLending)
	lending.POST("/", controllers.CreateLending)
	lending.PUT("/:id", controllers.UpdateLending)
	lending.DELETE("/:id", controllers.DeleteLending)
}
