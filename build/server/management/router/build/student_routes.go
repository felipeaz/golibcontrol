package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students/handler"
	"github.com/gin-gonic/gin"
)

// StudentRoutes initialize Category routes.
func StudentRoutes(rg *gin.RouterGroup, studentHandler handler.StudentHandler) {
	r := rg.Group("/students")
	r.GET("/", studentHandler.Get)
	r.GET("/:id", studentHandler.Find)
	r.POST("/", studentHandler.Create)
	r.PUT("/:id", studentHandler.Update)
	r.DELETE("/:id", studentHandler.Delete)
}
