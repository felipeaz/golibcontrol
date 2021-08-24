package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/platform/router/build"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	commentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/handler"
	reserveHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/handler"
	reviewHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/handler"
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

	cHandler := commentHandler.NewCommentHandler(dbService)
	build.CommentRoutes(vGroup, cHandler)

	resHandler := reserveHandler.NewReserveHandler(dbService)
	build.ReserveRoutes(vGroup, resHandler)

	revHandler := reviewHandler.NewReviewHandler(dbService)
	build.ReviewRoutes(vGroup, revHandler)

	return router.Run(":8083")
}
