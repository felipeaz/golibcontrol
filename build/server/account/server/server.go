package server

import (
	"github.com/FelipeAz/golibcontrol/build/router/account/router"
	"github.com/FelipeAz/golibcontrol/infra/middleware"
	authHandler "github.com/FelipeAz/golibcontrol/internal/app/account/auth/handler"
	authModule "github.com/FelipeAz/golibcontrol/internal/app/account/auth/module"
	"github.com/FelipeAz/golibcontrol/internal/app/account/users/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/account/users/module"
	"github.com/FelipeAz/golibcontrol/internal/app/account/users/repository"
	"github.com/FelipeAz/golibcontrol/internal/consumer"
	"github.com/FelipeAz/golibcontrol/internal/database"
	"github.com/FelipeAz/golibcontrol/internal/logger"
)

// Start initialize the webservice,
func Start(
	dbService database.GORMServiceInterface,
	cache database.Cache,
	consumersService consumer.Interface,
	log logger.LogInterface,
	mwr *middleware.Middleware) (err error) {
	accountRepository := repository.NewAccountRepository(dbService)
	accountModule := module.NewAccountModule(accountRepository, consumersService, cache, log)
	accountHandler := handler.NewAccountHandler(accountModule)
	aModule := authModule.NewAuthModule(accountRepository, consumersService, cache, log)
	aHandler := authHandler.NewAuthHandler(aModule)

	return router.Route(accountHandler, aHandler, mwr)
}
