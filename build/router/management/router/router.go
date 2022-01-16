package router

import (
	"github.com/FelipeAz/golibcontrol/build/router/management/router/routes"
	"github.com/FelipeAz/golibcontrol/infra/middleware"
	bookHandler "github.com/FelipeAz/golibcontrol/internal/app/management/books/handler"
	categoryHandler "github.com/FelipeAz/golibcontrol/internal/app/management/categories/handler"
	lendingHandler "github.com/FelipeAz/golibcontrol/internal/app/management/lending/handler"
	studentHandler "github.com/FelipeAz/golibcontrol/internal/app/management/students/handler"
	"github.com/gin-gonic/gin"
)

func Route(
	bHandler bookHandler.BookHandler,
	cHandler categoryHandler.CategoryHandler,
	sHandler studentHandler.StudentHandler,
	lHandler lendingHandler.LendingHandler,
	mwr *middleware.Middleware) error {
	router := gin.New()
	router.Use(mwr.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	routes.BookRoutes(vGroup, bHandler)
	routes.CategoryRoutes(vGroup, cHandler)
	routes.StudentRoutes(vGroup, sHandler)
	routes.LendingRoutes(vGroup, lHandler)

	return router.Run(":8081")
}
