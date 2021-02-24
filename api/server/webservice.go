package server

import (
	"github.com/FelipeAz/golibcontrol/api/server/router"
	"github.com/FelipeAz/golibcontrol/database"
)

// Start initialize the webservice
func Start() {
	db := database.Connect()
	defer database.CloseConnection(db)

	router.Run()
}
