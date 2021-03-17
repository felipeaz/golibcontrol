package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// StudentRoutes initialize Category routes.
func StudentRoutes(r *gin.Engine, studentHandler rest.StudentHandler) {
	rg := r.Group("/student")

	rg.GET("/", studentHandler.Get)
	rg.GET("/:id", studentHandler.Find)
	rg.POST("/", studentHandler.Create)
	rg.PUT("/:id", studentHandler.Update)
	rg.DELETE("/:id", studentHandler.Delete)
}
