package webservice

import (
	"github.com/FelipeAz/golibcontrol/database"
	"github.com/FelipeAz/golibcontrol/routes"
)

// Start initialize the webservice
func Start() {
	db := database.Connect()
	defer database.CloseConnection(db)

	routes.Run()
}
