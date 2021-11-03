package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/platform/router/build"
	commentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments/handler"
	replyHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/handler"
	reserveHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/handler"
	reviewHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func Build(
	cHandler commentHandler.CommentHandler,
	replHandler replyHandler.ReplyHandler,
	resHandler reserveHandler.ReserveHandler,
	revHandler reviewHandler.ReviewHandler,
) error {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.CommentRoutes(vGroup, cHandler)
	build.ReplyRoutes(vGroup, replHandler)
	build.ReserveRoutes(vGroup, resHandler)
	build.ReviewRoutes(vGroup, revHandler)

	return router.Run(":8083")
}
