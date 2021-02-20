package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func getRoutes() {
	r := router.Group("/")

	addPingRoute(r)
	addStudentRoutes(r)
	addBookRoutes(r)
	addCategoryRoute(r)
	addLendingRoutes(r)
}

// Run Starts the server
func Run() {
	getRoutes()
	router.Run(":8080")
}
