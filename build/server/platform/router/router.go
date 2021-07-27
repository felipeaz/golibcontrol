package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/platform/router/build"
	"github.com/FelipeAz/golibcontrol/infra/jwt"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/infra/redis"
	commentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/handler"
	reserveHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/handler"
	reviewHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/handler"

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

	cHandler := commentHandler.NewCommentHandler(dbService)
	build.CommentRoutes(tokenAuthMiddleware, vGroup, cHandler)

	resHandler := reserveHandler.NewReserveHandler(dbService)
	build.ReserveRoutes(tokenAuthMiddleware, vGroup, resHandler)

	revHandler := reviewHandler.NewReviewHandler(dbService)
	build.ReviewRoutes(tokenAuthMiddleware, vGroup, revHandler)

	return router.Run(":8083")
}
