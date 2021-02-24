package router

import (
	"github.com/FelipeAz/golibcontrol/controllers"
	"github.com/gin-gonic/gin"
)

// AddBookRoutes
func addBookRoutes(r *gin.RouterGroup) {
	book := r.Group("book")

	book.GET("/", controllers.GetBooks)
	book.GET("/:id", controllers.GetBook)
	book.POST("/", controllers.CreateBook)
	book.PUT("/:id", controllers.UpdateBook)
	book.DELETE("/:id", controllers.DeleteBook)
}
