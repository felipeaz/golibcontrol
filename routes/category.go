package routes

import (
	"github.com/FelipeAz/golibcontrol/controllers"
	"github.com/gin-gonic/gin"
)

func addCategoryRoute(r *gin.RouterGroup) {
	category := r.Group("category")

	category.GET("/", controllers.GetCategories)
	category.GET("/:id", controllers.GetCategory)
	category.POST("/", controllers.CreateCategory)
	category.PUT("/:id", controllers.UpdateCategory)
	category.DELETE("/:id", controllers.DeleteCategory)
}
