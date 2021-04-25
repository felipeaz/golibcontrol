package router

import (
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/FelipeAz/golibcontrol/platform/jwt"
	"github.com/FelipeAz/golibcontrol/platform/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
	"github.com/FelipeAz/golibcontrol/platform/router/build"
)

func buildRoutes(db *gorm.DB, cache *redis.Cache) error {
	router := gin.Default()
	jwtAuth := jwt.NewAuth(cache)
	tokenAuthMiddleware := middleware.NewTokenMiddleware(jwtAuth)

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	accountHandler := rest.NewAccountHandler(db, cache)
	build.AccountRoutes(vGroup, accountHandler)

	bookHandler := rest.NewBookHandler(db)
	build.BookRoutes(tokenAuthMiddleware, vGroup, bookHandler)

	categoryHandler := rest.NewCategoryHandler(db)
	build.CategoryRoutes(tokenAuthMiddleware, vGroup, categoryHandler)

	studentHandler := rest.NewStudentHandler(db)
	build.StudentRoutes(tokenAuthMiddleware, vGroup, studentHandler)

	lendingHandler := rest.NewLendingHandler(db)
	build.LendingRoutes(tokenAuthMiddleware, vGroup, lendingHandler)

	return router.Run()
}

// Run Starts the server
func Run(db *gorm.DB, cache *redis.Cache) error {
	return buildRoutes(db, cache)
}
