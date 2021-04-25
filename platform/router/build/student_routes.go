package build

import (
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
)

// StudentRoutes initialize Category routes.
func StudentRoutes(middleware *middleware.TokenMiddleware, rg *gin.RouterGroup, studentHandler rest.StudentHandler) {
	r := rg.Group("/student")

	r.GET("/", middleware.TokenAuth(), studentHandler.Get)
	r.GET("/:id", middleware.TokenAuth(), studentHandler.Find)
	r.POST("/", middleware.TokenAuth(), studentHandler.Create)
	r.PUT("/:id", middleware.TokenAuth(), studentHandler.Update)
	r.DELETE("/:id", middleware.TokenAuth(), studentHandler.Delete)
}
