package server

import (
	"log"

	"github.com/FelipeAz/golibcontrol/build/server/management/router"
	"github.com/FelipeAz/golibcontrol/infra/mysql/management/database"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
)

// Start initialize the webservice,
func Start(user, password, host, port, databaseName string) (err error) {
	db, err := database.Connect(user, password, host, port, databaseName)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer database.CloseConnection(db)

	dbService, err := service.NewMySQLService(db)
	return router.Run(dbService)
}
