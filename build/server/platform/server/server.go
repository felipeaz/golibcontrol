package server

import (
	"github.com/FelipeAz/golibcontrol/build/router/platform/router"
	commentHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/comments/handler"
	commentModule "github.com/FelipeAz/golibcontrol/internal/app/platform/comments/module"
	commentRepository "github.com/FelipeAz/golibcontrol/internal/app/platform/comments/repository"
	conferenceHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/conferences/handler"
	conferenceModule "github.com/FelipeAz/golibcontrol/internal/app/platform/conferences/module"
	conferenceRepository "github.com/FelipeAz/golibcontrol/internal/app/platform/conferences/repository"
	groupHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/groups/handler"
	groupModule "github.com/FelipeAz/golibcontrol/internal/app/platform/groups/module"
	groupRepository "github.com/FelipeAz/golibcontrol/internal/app/platform/groups/repository"
	replyHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/replies/handler"
	replyModule "github.com/FelipeAz/golibcontrol/internal/app/platform/replies/module"
	replyRepository "github.com/FelipeAz/golibcontrol/internal/app/platform/replies/repository"
	reserveHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/reserves/handler"
	reserveModule "github.com/FelipeAz/golibcontrol/internal/app/platform/reserves/module"
	reserveRepository "github.com/FelipeAz/golibcontrol/internal/app/platform/reserves/repository"
	reviewHandler "github.com/FelipeAz/golibcontrol/internal/app/platform/reviews/handler"
	reviewModule "github.com/FelipeAz/golibcontrol/internal/app/platform/reviews/module"
	reviewRepository "github.com/FelipeAz/golibcontrol/internal/app/platform/reviews/repository"
	"github.com/FelipeAz/golibcontrol/internal/database"
	"github.com/FelipeAz/golibcontrol/internal/logger"
)

// Start initialize the webservice,
func Start(dbService database.GORMServiceInterface, log logger.LogInterface) (err error) {
	cRepository := commentRepository.NewCommentRepository(dbService)
	cModule := commentModule.NewCommentModule(cRepository, log)
	cHandler := commentHandler.NewCommentHandler(cModule)

	replRepository := replyRepository.NewReplyRepository(dbService)
	replModule := replyModule.NewReplyModule(replRepository, log)
	replHandler := replyHandler.NewReplyHandler(replModule)

	resRepository := reserveRepository.NewReserveRepository(dbService)
	resModule := reserveModule.NewReserveModule(resRepository, log)
	resHandler := reserveHandler.NewReserveHandler(resModule)

	revRepository := reviewRepository.NewReviewRepository(dbService)
	revModule := reviewModule.NewReviewModule(revRepository, log)
	revHandler := reviewHandler.NewReviewHandler(revModule)

	confRepository := conferenceRepository.NewConferenceRepository(dbService)
	confModule := conferenceModule.NewConferenceModule(confRepository, log)
	confHandler := conferenceHandler.NewConferenceHandler(confModule)

	grpRepository := groupRepository.NewGroupRepository(dbService)
	grpModule := groupModule.NewGroupModule(grpRepository, log)
	grpHandler := groupHandler.NewGroupHandler(grpModule)

	return router.Route(router.Handlers{
		CommentHandler:    cHandler,
		ReplyHandler:      replHandler,
		ReserveHandler:    resHandler,
		ReviewHandler:     revHandler,
		ConferenceHandler: confHandler,
		GroupHandler:      grpHandler,
	})
}
