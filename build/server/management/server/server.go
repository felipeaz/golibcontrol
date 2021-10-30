package server

import (
	"github.com/FelipeAz/golibcontrol/build/server/management/router"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	bookHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/handler"
	bookModule "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/module"
	bookRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/repository"
	categoryHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/handler"
	categoryModule "github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/module"
	categoryRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/management/category/repository"
	lendingHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/handler"
	lendingModule "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/module"
	lendingRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/repository"
	studentHandler "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/handler"
	studentModule "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/module"
	studentRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/repository"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// Start initialize the webservice,
func Start(dbService database.GORMServiceInterface, log logger.LogInterface) (err error) {
	bcRepository := bookRepository.NewBookCategoryRepository(dbService)
	bRepository := bookRepository.NewBookRepository(dbService, bcRepository)
	bModule := bookModule.NewBookModule(bRepository, log)
	bHandler := bookHandler.NewBookHandler(bModule)

	cRepository := categoryRepository.NewCategoryRepository(dbService)
	cModule := categoryModule.NewCategoryModule(cRepository, log)
	cHandler := categoryHandler.NewCategoryHandler(cModule)

	sRepository := studentRepository.NewStudentRepository(dbService)
	sModule := studentModule.NewStudentModule(sRepository, log)
	sHandler := studentHandler.NewStudentHandler(sModule)

	lRepository := lendingRepository.NewLendingRepository(dbService, sRepository, bRepository)
	lModule := lendingModule.NewLendingModule(lRepository, log)
	lHandler := lendingHandler.NewLendingHandler(lModule)

	return router.Build(bHandler, cHandler, sHandler, lHandler)
}
