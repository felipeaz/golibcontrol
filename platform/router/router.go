package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/FelipeAz/golibcontrol/internal/app/handler"
	"github.com/FelipeAz/golibcontrol/platform/router/build"
)

func buildRoutes(db *gorm.DB) (err error) {
	router := gin.Default()

	bookHandler := handler.NewBookHandler(db)
	build.BookRoutes(router, bookHandler)

	//categoryHandler := handler.NewCategoryHandler(db)
	//build.CategoryRoute(router)

	//studentHandler := handler.NewStudentHandler(db)
	//build.StudentRoutes(router)

	//lendingHandler := handler.NewLendingHandler
	//build.LendingRoutes(router)

	err = router.Run()
	return
}

// Run Starts the server
func Run(db *gorm.DB) (err error) {
	err = buildRoutes(db)
	return
}
