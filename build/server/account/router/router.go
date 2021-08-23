package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/account/router/build"
	"github.com/FelipeAz/golibcontrol/infra/auth"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/infra/redis"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// Run Starts the server
func Run(dbService *service.MySQLService, apiGatewayAuth auth.Auth, cache *redis.Cache) error {
	return buildRoutes(dbService, apiGatewayAuth, cache)
}

func buildRoutes(dbService *service.MySQLService, apiGatewayAuth auth.Auth, cache *redis.Cache) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	aHandler := handler.NewAccountHandler(dbService, apiGatewayAuth, cache)
	build.UserRoutes(vGroup, aHandler)

	return router.Run(":8082")
}
