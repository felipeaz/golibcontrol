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

	categoryHandler := handler.NewCategoryHandler(db)
	build.CategoryRoutes(router, categoryHandler)

	studentHandler := handler.NewStudentHandler(db)
	build.StudentRoutes(router, studentHandler)

	lendingHandler := handler.NewLendingHandler(db)
	build.LendingRoutes(router, lendingHandler)

	err = router.Run()
	return
}

// Run Starts the server
func Run(db *gorm.DB) (err error) {
	err = buildRoutes(db)
	return
}
