package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/platform/router/build"
	commentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/handler"
	reserveHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/handler"
	reviewHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func Build(
	cHandler commentHandler.CommentHandler,
	resHandler reserveHandler.ReserveHandler,
	revHandler reviewHandler.ReviewHandler,
) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.CommentsRoutes(vGroup, cHandler)
	build.ReservesRoutes(vGroup, resHandler)
	build.ReviewsRoutes(vGroup, revHandler)

	return router.Run(":8083")
}
