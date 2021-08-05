package server

import (
	"log"

	"github.com/FelipeAz/golibcontrol/build/server/account/router"
	"github.com/FelipeAz/golibcontrol/infra/auth"
	"github.com/FelipeAz/golibcontrol/infra/mysql/account/database"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/internal/pkg/http/request"
)

// Start initialize the webservice,
func Start(user, password, host, port, databaseName, consumersHost string) (err error) {
	db, err := database.Connect(user, password, host, port, databaseName)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer database.CloseConnection(db)

	dbService, err := service.NewMySQLService(db)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	apiGatewayAuth := auth.NewAuth(request.NewHttpRequest(consumersHost))
	return router.Run(dbService, apiGatewayAuth)
}
