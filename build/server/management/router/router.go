package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/management/router/build"
	bookHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/handler"
	categoryHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/handler"
	lendingHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/handler"
	studentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func Build(
	bHandler bookHandler.BookHandler,
	cHandler categoryHandler.CategoryHandler,
	sHandler studentHandler.StudentHandler,
	lHandler lendingHandler.LendingHandler,
) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.BookRoutes(vGroup, bHandler)
	build.CategoryRoutes(vGroup, cHandler)
	build.StudentRoutes(vGroup, sHandler)
	build.LendingRoutes(vGroup, lHandler)

	return router.Run(":8081")
}
