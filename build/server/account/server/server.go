package server

import (
	"github.com/FelipeAz/golibcontrol/build/server/account/router"
	"github.com/FelipeAz/golibcontrol/internal/app/auth"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/module"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/repository"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// Start initialize the webservice,
func Start(
	dbService database.GORMServiceInterface,
	cache database.CacheInterface,
	apiGatewayAuth auth.AuthInterface,
	log logger.LogInterface,
) (err error) {
	accountRepository := repository.NewAccountRepository(dbService)
	accountModule := module.NewAccountModule(accountRepository, apiGatewayAuth, cache, log)
	accountHandler := handler.NewAccountHandler(accountModule)
	return router.Build(accountHandler)
}
