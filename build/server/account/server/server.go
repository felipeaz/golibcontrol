package server

import (
	"github.com/FelipeAz/golibcontrol/build/router/account/router"
	authHandler "github.com/FelipeAz/golibcontrol/internal/app/account/auth/handler"
	authModule "github.com/FelipeAz/golibcontrol/internal/app/account/auth/module"
	accountHandler "github.com/FelipeAz/golibcontrol/internal/app/account/users/handler"
	accountModule "github.com/FelipeAz/golibcontrol/internal/app/account/users/module"
	accountRepository "github.com/FelipeAz/golibcontrol/internal/app/account/users/repository"
	"github.com/FelipeAz/golibcontrol/internal/consumer"
	"github.com/FelipeAz/golibcontrol/internal/database"
	"github.com/FelipeAz/golibcontrol/internal/logger"
)

// Start initialize the webservice,
func Start(dbService database.GORMServiceInterface, cache database.Cache, consumersService consumer.Interface, log logger.LogInterface) (err error) {
	accRepository := accountRepository.NewAccountRepository(dbService)
	accModule := accountModule.NewAccountModule(accRepository, consumersService, cache, log)
	accHandler := accountHandler.NewAccountHandler(accModule)
	aModule := authModule.NewAuthModule(accRepository, consumersService, cache, log)
	aHandler := authHandler.NewAuthHandler(aModule)

	return router.Route(router.Handlers{
		AccountHandler: accHandler,
		AuthHandler:    aHandler,
	})
}
