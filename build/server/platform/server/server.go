package server

import (
	"github.com/FelipeAz/golibcontrol/build/server/platform/router"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	commentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/handler"
	commentModule "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/module"
	commentRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/repository"
	reserveHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/handler"
	reserveModule "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/module"
	reserveRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/repository"
	reviewHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/handler"
	reviewModule "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/module"
	reviewRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/repository"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// Start initialize the webservice,
func Start(dbService database.GORMServiceInterface, log logger.LogInterface) (err error) {
	cRepository := commentRepository.NewCommentRepository(dbService)
	cModule := commentModule.NewCommentModule(cRepository, log)
	cHandler := commentHandler.NewCommentHandler(cModule)

	resRepository := reserveRepository.NewReserveRepository(dbService)
	resModule := reserveModule.NewReserveModule(resRepository, log)
	resHandler := reserveHandler.NewReserveHandler(resModule)

	revRepository := reviewRepository.NewReviewRepository(dbService)
	revModule := reviewModule.NewReviewModule(revRepository, log)
	revHandler := reviewHandler.NewReviewHandler(revModule)

	return router.Build(cHandler, resHandler, revHandler)
}
