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

type Handlers struct {
	CommentHandler    commentHandler.CommentHandler
	ReplyHandler      replyHandler.ReplyHandler
	ReserveHandler    reserveHandler.ReserveHandler
	ReviewHandler     reviewHandler.ReviewHandler
	ConferenceHandler conferenceHandler.ConferenceHandler
	GroupHandler      groupHandler.GroupHandler
}

func Route(handlers Handlers) error {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())

	apiRg := router.Group("/api")
	vGroup := apiRg.Group("/v1")

	routes.CommentRoutes(vGroup, handlers.CommentHandler)
	routes.ReplyRoutes(vGroup, handlers.ReplyHandler)
	routes.ReserveRoutes(vGroup, handlers.ReserveHandler)
	routes.ReviewRoutes(vGroup, handlers.ReviewHandler)
	routes.ConferenceRoutes(vGroup, handlers.ConferenceHandler)
	routes.GroupRoutes(vGroup, handlers.GroupHandler)

	return router.Run(":8083")
}
