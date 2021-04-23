package router

import (
	"github.com/FelipeAz/golibcontrol/platform/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/handler/rest"
	"github.com/FelipeAz/golibcontrol/platform/router/build"
)

func buildRoutes(db *gorm.DB, cache *redis.Cache) error {
	router := gin.Default()

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	bookHandler := rest.NewBookHandler(db)
	build.BookRoutes(vGroup, bookHandler)

	categoryHandler := rest.NewCategoryHandler(db)
	build.CategoryRoutes(vGroup, categoryHandler)

	studentHandler := rest.NewStudentHandler(db)
	build.StudentRoutes(vGroup, studentHandler)

	lendingHandler := rest.NewLendingHandler(db)
	build.LendingRoutes(vGroup, lendingHandler)

	accountHandler := rest.NewAccountHandler(db, cache)
	build.AccountRoutes(vGroup, accountHandler)

	return router.Run()
}

// Run Starts the server
func Run(db *gorm.DB, cache *redis.Cache) error {
	return buildRoutes(db, cache)
}
