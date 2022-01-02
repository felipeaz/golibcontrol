package router

import (
	"github.com/FelipeAz/golibcontrol/build/server/platform/router/build"
	"github.com/FelipeAz/golibcontrol/internal/app/middleware"
	commentHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/comments/handler"
	conferenceHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/conferences/handler"
	groupHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/groups/handler"
	replyHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/replies/handler"
	reserveHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/reserves/handler"
	reviewHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/reviews/handler"

	"github.com/gin-gonic/gin"
)

func Build(
	cHandler commentHandler.CommentHandler,
	replHandler replyHandler.ReplyHandler,
	resHandler reserveHandler.ReserveHandler,
	revHandler reviewHandler.ReviewHandler,
	confHandler conferenceHandler.ConferenceHandler,
	grpHandler groupHandler.GroupHandler,
) error {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	build.CommentRoutes(vGroup, cHandler)
	build.ReplyRoutes(vGroup, replHandler)
	build.ReserveRoutes(vGroup, resHandler)
	build.ReviewRoutes(vGroup, revHandler)
	build.ConferenceRoutes(vGroup, confHandler)
	build.GroupRoutes(vGroup, grpHandler)

	return router.Run(":8083")
}
