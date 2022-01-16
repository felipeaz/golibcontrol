package router

import (
	"github.com/FelipeAz/golibcontrol/build/router/platform/router/routes"
	"github.com/FelipeAz/golibcontrol/infra/middleware"
	commentHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/comments/handler"
	conferenceHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/conferences/handler"
	groupHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/groups/handler"
	replyHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/replies/handler"
	reserveHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/reserves/handler"
	reviewHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/reviews/handler"

	"github.com/gin-gonic/gin"
)

func Route(
	cHandler commentHandler.CommentHandler,
	replHandler replyHandler.ReplyHandler,
	resHandler reserveHandler.ReserveHandler,
	revHandler reviewHandler.ReviewHandler,
	confHandler conferenceHandler.ConferenceHandler,
	grpHandler groupHandler.GroupHandler,
	mwr *middleware.Middleware) error {
	router := gin.New()
	router.Use(mwr.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	routes.CommentRoutes(vGroup, cHandler)
	routes.ReplyRoutes(vGroup, replHandler)
	routes.ReserveRoutes(vGroup, resHandler)
	routes.ReviewRoutes(vGroup, revHandler)
	routes.ConferenceRoutes(vGroup, confHandler)
	routes.GroupRoutes(vGroup, grpHandler)

	return router.Run(":8083")
}
