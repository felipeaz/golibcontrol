package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler"
)

// StudentRoutes initialize Category routes.
func StudentRoutes(r *gin.Engine, studentHandler handler.StudentHandler) {
	rg := r.Group("/student")

	rg.GET("/", studentHandler.Get)
	rg.GET("/:id", studentHandler.Find)
	rg.POST("/", studentHandler.Create)
	rg.PUT("/:id", studentHandler.Update)
	rg.DELETE("/:id", studentHandler.Delete)
}
