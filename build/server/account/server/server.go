package server

import (
	"github.com/FelipeAz/golibcontrol/build/server/account/router"
	auth_handler "github.com/FelipeAz/golibcontrol/internal/app/account/auth/handler"
	auth_module "github.com/FelipeAz/golibcontrol/internal/app/account/auth/module"
	"github.com/FelipeAz/golibcontrol/internal/app/account/users/handler"
	"github.com/FelipeAz/golibcontrol/internal/app/account/users/module"
	"github.com/FelipeAz/golibcontrol/internal/app/account/users/repository"
	"github.com/FelipeAz/golibcontrol/internal/app/consumer"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

// Start initialize the webservice,
func Start(
	dbService database.GORMServiceInterface,
	cache database.Cache,
	consumersService consumer.Interface,
	log logger.LogInterface,
) (err error) {
	accountRepository := repository.NewAccountRepository(dbService)
	accountModule := module.NewAccountModule(accountRepository, consumersService, cache, log)
	accountHandler := handler.NewAccountHandler(accountModule)
	authModule := auth_module.NewAuthModule(accountRepository, consumersService, cache, log)
	authHandler := auth_handler.NewAuthHandler(authModule)
	return router.Build(accountHandler, authHandler)
}
