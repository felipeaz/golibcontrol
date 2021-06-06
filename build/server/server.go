package server

import (
	"log"

	"github.com/FelipeAz/golibcontrol/platform/mysql"
	"github.com/FelipeAz/golibcontrol/platform/mysql/service"
	"github.com/FelipeAz/golibcontrol/platform/redis"
	"github.com/FelipeAz/golibcontrol/platform/router"
)

// Start initialize the webservice
func Start() (err error) {
	dbService, err := service.NewMySQLService()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer mysql.CloseConnection(dbService.DB)

	cache := redis.NewCache()

	return router.Run(dbService, cache)
}
