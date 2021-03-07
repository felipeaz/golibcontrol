package router

import (
	"github.com/FelipeAz/golibcontrol/platform/router/build"
	"github.com/gin-gonic/gin"
)

func buildRoutes() (err error) {
	router := gin.Default()

	build.StudentRoutes(router)
	build.BookRoutes(router)
	build.CategoryRoute(router)
	build.LendingRoutes(router)

	err = router.Run()
	return
}

// Run Starts the server
func Run() (err error) {
	err = buildRoutes()
	return
}
