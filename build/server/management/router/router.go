package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/management/router/build"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	bookHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/handler"
	categoryHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/handler"
	lendingHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/handler"
	studentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

// Run Starts the server
func Run(dbService database.GORMServiceInterface) error {
	return buildRoutes(dbService)
}

func buildRoutes(dbService database.GORMServiceInterface) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	bHandler := bookHandler.NewBookHandler(dbService)
	build.BookRoutes(vGroup, bHandler)

	cHandler := categoryHandler.NewCategoryHandler(dbService)
	build.CategoryRoutes(vGroup, cHandler)

	sHandler := studentHandler.NewStudentHandler(dbService)
	build.StudentRoutes(vGroup, sHandler)

	lHandler := lendingHandler.NewLendingHandler(dbService)
	build.LendingRoutes(vGroup, lHandler)

	return router.Run(":8081")
}
