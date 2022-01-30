package routes

import (
	"github.com/FelipeAz/golibcontrol/internal/app/management/registries/handler"
	"github.com/gin-gonic/gin"
)

// RegistryRoutes initialize Book routes.
func RegistryRoutes(rg *gin.RouterGroup, registryHander handler.RegistryHandler) {
	r := rg.Group("/registries")
	r.GET("", registryHander.Get)
	r.GET("/:id", registryHander.Find)
	r.POST("", registryHander.Create)
	r.PUT("/:id", registryHander.Update)
	r.DELETE("/:id", registryHander.Delete)
}
