package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/management/router/build"
	"github.com/FelipeAz/golibcontrol/infra/jwt"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/infra/redis"
	bookHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/handler"
	categoryHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/handler"
	lendingHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/handler"
	studentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/handler"
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

	bHandler := bookHandler.NewBookHandler(dbService)
	build.BookRoutes(tokenAuthMiddleware, vGroup, bHandler)

	cHandler := categoryHandler.NewCategoryHandler(dbService)
	build.CategoryRoutes(tokenAuthMiddleware, vGroup, cHandler)

	sHandler := studentHandler.NewStudentHandler(dbService)
	build.StudentRoutes(tokenAuthMiddleware, vGroup, sHandler)

	lHandler := lendingHandler.NewLendingHandler(dbService)
	build.LendingRoutes(tokenAuthMiddleware, vGroup, lHandler)

	return router.Run(":8081")
}
