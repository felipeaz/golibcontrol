package build

import (
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// StudentRoutes initialize Category routes.
func StudentRoutes(rg *gin.RouterGroup, studentHandler rest.StudentHandler) {
	r := rg.Group("/student")

	r.GET("/", studentHandler.Get)
	r.GET("/:id", studentHandler.Find)
	r.POST("/", studentHandler.Create)
	r.PUT("/:id", studentHandler.Update)
	r.DELETE("/:id", studentHandler.Delete)
}
