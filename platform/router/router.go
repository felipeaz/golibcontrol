package router

import (
	account_handler "github.com/FelipeAz/golibcontrol/internal/app/domain/account/handler"
	book_handler "github.com/FelipeAz/golibcontrol/internal/app/domain/book/handler"
	category_handler "github.com/FelipeAz/golibcontrol/internal/app/domain/category/handler"
	lending_handler "github.com/FelipeAz/golibcontrol/internal/app/domain/lending/handler"
	student_handler "github.com/FelipeAz/golibcontrol/internal/app/domain/student/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/FelipeAz/golibcontrol/platform/jwt"
	"github.com/FelipeAz/golibcontrol/platform/mysql/service"
	"github.com/FelipeAz/golibcontrol/platform/redis"
	"github.com/FelipeAz/golibcontrol/platform/router/build"
	"github.com/gin-gonic/gin"
)

func buildRoutes(dbService *service.MySQLService, cache *redis.Cache) error {
	router := gin.Default()
	jwtAuth := jwt.NewAuth(cache)
	tokenAuthMiddleware := middleware.NewTokenMiddleware(jwtAuth)

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	accountHandler := account_handler.NewAccountHandler(jwtAuth, dbService, cache)
	build.AccountRoutes(tokenAuthMiddleware, vGroup, accountHandler)

	bookHandler := book_handler.NewBookHandler(dbService)
	build.BookRoutes(tokenAuthMiddleware, vGroup, bookHandler)

	categoryHandler := category_handler.NewCategoryHandler(dbService)
	build.CategoryRoutes(tokenAuthMiddleware, vGroup, categoryHandler)

	studentHandler := student_handler.NewStudentHandler(dbService)
	build.StudentRoutes(tokenAuthMiddleware, vGroup, studentHandler)

	lendingHandler := lending_handler.NewLendingHandler(dbService)
	build.LendingRoutes(tokenAuthMiddleware, vGroup, lendingHandler)

	return router.Run()
}

// Run Starts the server
func Run(dbService *service.MySQLService, cache *redis.Cache) error {
	return buildRoutes(dbService, cache)
}
