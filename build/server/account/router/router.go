package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/account/router/build"
	"github.com/FelipeAz/golibcontrol/infra/auth"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/handler"
	"github.com/gin-gonic/gin"
)

// Run Starts the server
func Run(dbService *service.MySQLService, apiGatewayAuth auth.Auth) error {
	return buildRoutes(dbService, apiGatewayAuth)
}

func buildRoutes(dbService *service.MySQLService, apiGatewayAuth auth.Auth) error {
	router := gin.Default()

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	aHandler := handler.NewAccountHandler(dbService, apiGatewayAuth)
	build.UserRoutes(vGroup, aHandler)

	return router.Run(":8082")
}
