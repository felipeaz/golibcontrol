package server

import (
	"github.com/FelipeAz/golibcontrol/build/router/management/router"
	"github.com/FelipeAz/golibcontrol/build/server/management/grpc"
	bookHandler "github.com/FelipeAz/golibcontrol/internal/app/management/books/handler"
	bookModule "github.com/FelipeAz/golibcontrol/internal/app/management/books/module"
	"github.com/FelipeAz/golibcontrol/internal/app/management/books/repository"
	categoryHandler "github.com/FelipeAz/golibcontrol/internal/app/management/categories/handler"
	categoryModule "github.com/FelipeAz/golibcontrol/internal/app/management/categories/module"
	categoryRepository "github.com/FelipeAz/golibcontrol/internal/app/management/categories/repository"
	lendingHandler "github.com/FelipeAz/golibcontrol/internal/app/management/lending/handler"
	lendingModule "github.com/FelipeAz/golibcontrol/internal/app/management/lending/module"
	lendingRepository "github.com/FelipeAz/golibcontrol/internal/app/management/lending/repository"
	registryHandler "github.com/FelipeAz/golibcontrol/internal/app/management/registries/handler"
	registryModule "github.com/FelipeAz/golibcontrol/internal/app/management/registries/module"
	registryRepository "github.com/FelipeAz/golibcontrol/internal/app/management/registries/repository"
	studentHandler "github.com/FelipeAz/golibcontrol/internal/app/management/students/handler"
	studentModule "github.com/FelipeAz/golibcontrol/internal/app/management/students/module"
	studentRepository "github.com/FelipeAz/golibcontrol/internal/app/management/students/repository"
	"github.com/FelipeAz/golibcontrol/internal/database"
	"github.com/FelipeAz/golibcontrol/internal/logger"
)

// Start initialize the webservice,
func Start(dbService database.GORMServiceInterface, log logger.LogInterface) (err error) {
	bcRepository := repository.NewBookCategoryRepository(dbService)
	bRepository := repository.NewBookRepository(dbService)
	bModule := bookModule.NewBookModule(bRepository, bcRepository, log)
	bHandler := bookHandler.NewBookHandler(bModule)

	cRepository := categoryRepository.NewCategoryRepository(dbService)
	cModule := categoryModule.NewCategoryModule(cRepository, log)
	cHandler := categoryHandler.NewCategoryHandler(cModule)

	rRepository := registryRepository.NewRegistryRepository(dbService)
	rModule := registryModule.NewRegistryModule(rRepository, log)
	rHandler := registryHandler.NewRegistryHandler(rModule)

	sRepository := studentRepository.NewStudentRepository(dbService)
	sModule := studentModule.NewStudentModule(sRepository, log)
	sHandler := studentHandler.NewStudentHandler(sModule)

	lRepository := lendingRepository.NewLendingRepository(dbService)
	lModule := lendingModule.NewLendingModule(lRepository, log)
	lHandler := lendingHandler.NewLendingHandler(lModule)

	go func() {
		err = grpc.Start(rModule)
		if err != nil {
			panic(err)
		}
	}()

	return router.Route(router.Handlers{
		BookHandler:     bHandler,
		CategoryHandler: cHandler,
		StudentHandler:  sHandler,
		LendingHandler:  lHandler,
		RegistryHandler: rHandler,
	})
}
