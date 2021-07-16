package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/account/router/build"
	"github.com/FelipeAz/golibcontrol/infra/jwt"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/infra/redis"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// Run Starts the server
func Run(dbService *service.MySQLService, cache *redis.Cache) error {
	return buildRoutes(dbService, cache)
}

func buildRoutes(dbService *service.MySQLService, cache *redis.Cache) error {
	router := gin.Default()
	jwtAuth := jwt.NewAuth(cache)
	tokenAuthMiddleware := middleware.NewTokenMiddleware(jwtAuth)

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	aHandler := handler.NewAccountHandler(jwtAuth, dbService, cache)
	build.UserRoutes(tokenAuthMiddleware, vGroup, aHandler)

	return router.Run(":8082")
}
