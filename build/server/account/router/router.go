package router

import (
	"os"

	"github.com/FelipeAz/golibcontrol/build/server/account/router/build"
	"github.com/FelipeAz/golibcontrol/internal/app/auth"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// Run Starts the server
func Run(
	dbService database.GORMServiceInterface,
	apiGatewayAuth auth.AuthInterface,
	cache database.CacheInterface) error {
	return buildRoutes(dbService, apiGatewayAuth, cache)
}

func buildRoutes(
	dbService database.GORMServiceInterface,
	apiGatewayAuth auth.AuthInterface,
	cache database.CacheInterface) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	aHandler := handler.NewAccountHandler(dbService, apiGatewayAuth, cache, os.Getenv("JWT_SECRET_KEY"))
	build.UserRoutes(vGroup, aHandler)

	return router.Run(":8082")
}
