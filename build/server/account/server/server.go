package server

import (
	"log"

	"github.com/FelipeAz/golibcontrol/build/server/account/router"
	"github.com/FelipeAz/golibcontrol/infra/auth"
	"github.com/FelipeAz/golibcontrol/infra/auth/http/request"
	"github.com/FelipeAz/golibcontrol/infra/logger"
	"github.com/FelipeAz/golibcontrol/infra/mysql/account/database"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/infra/redis"
)

// Start initialize the webservice,
func Start(user, password, host, port, databaseName, consumersHost, cacheHost, cachePort, cacheExpireTime string) (err error) {
	db, err := database.Connect(user, password, host, port, databaseName)
	if err != nil {
		logger.LogError(err)
		log.Fatal(err.Error())
	}
	defer database.CloseConnection(db)

	dbService, err := service.NewMySQLService(db)
	if err != nil {
		logger.LogError(err)
		log.Fatal(err.Error())
	}

	cache := redis.NewCache(cacheHost, cachePort, cacheExpireTime)

	apiGatewayAuth := auth.NewAuth(request.NewHttpRequest(consumersHost))
	return router.Run(dbService, apiGatewayAuth, cache)
}
