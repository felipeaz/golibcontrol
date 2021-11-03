package server

import (
	"github.com/FelipeAz/golibcontrol/build/server/platform/router"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	commentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments/handler"
	commentModule "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments/module"
	commentRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comments/repository"
	replyHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/handler"
	replyModule "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/module"
	replyRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies/repository"
	reserveHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/handler"
	reserveModule "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/module"
	reserveRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/repository"
	reviewHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews/handler"
	reviewModule "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews/module"
	reviewRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews/repository"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
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

	return router.Build(cHandler, replHandler, resHandler, revHandler)
}
