package server

import (
	"log"

	"github.com/FelipeAz/golibcontrol/build/server/management/router"
	"github.com/FelipeAz/golibcontrol/infra/mysql/management/database"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/infra/redis"
)

// Start initialize the webservice,
func Start() (err error) {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer database.CloseConnection(db)

	dbService, err := service.NewMySQLService(db)
	cache := redis.NewCache()
	return router.Run(dbService, cache)
}
