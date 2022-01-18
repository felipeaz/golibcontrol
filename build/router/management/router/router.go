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

type Handlers struct {
	BookHandler     bookHandler.BookHandler
	CategoryHandler categoryHandler.CategoryHandler
	StudentHandler  studentHandler.StudentHandler
	LendingHandler  lendingHandler.LendingHandler
}

func Route(handlers Handlers) error {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	routes.BookRoutes(vGroup, handlers.BookHandler)
	routes.CategoryRoutes(vGroup, handlers.CategoryHandler)
	routes.StudentRoutes(vGroup, handlers.StudentHandler)
	routes.LendingRoutes(vGroup, handlers.LendingHandler)

	return router.Run(":8081")
}
