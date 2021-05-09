package router

import (
	accountHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/account/handler"
	bookHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/book/handler"
	categoryHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/category/handler"
	lendingHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/lending/handler"
	studentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/student/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/FelipeAz/golibcontrol/platform/jwt"
	"github.com/FelipeAz/golibcontrol/platform/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/platform/router/build"
)

func buildRoutes(db *gorm.DB, cache *redis.Cache) error {
	router := gin.Default()
	jwtAuth := jwt.NewAuth(cache)
	tokenAuthMiddleware := middleware.NewTokenMiddleware(jwtAuth)

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	accountHandler := accountHandler.NewAccountHandler(jwtAuth, db, cache)
	build.AccountRoutes(tokenAuthMiddleware, vGroup, accountHandler)

	bookHandler := bookHandler.NewBookHandler(db)
	build.BookRoutes(tokenAuthMiddleware, vGroup, bookHandler)

	categoryHandler := categoryHandler.NewCategoryHandler(db)
	build.CategoryRoutes(tokenAuthMiddleware, vGroup, categoryHandler)

	studentHandler := studentHandler.NewStudentHandler(db)
	build.StudentRoutes(tokenAuthMiddleware, vGroup, studentHandler)

	lendingHandler := lendingHandler.NewLendingHandler(db)
	build.LendingRoutes(tokenAuthMiddleware, vGroup, lendingHandler)

	return router.Run()
}

// Run Starts the server
func Run(db *gorm.DB, cache *redis.Cache) error {
	return buildRoutes(db, cache)
}
