package server

import (
	"log"

	"github.com/FelipeAz/golibcontrol/platform/mysql"
	"github.com/FelipeAz/golibcontrol/platform/router"
)

// Start initialize the webservice
func Start() (err error) {
	db, err := mysql.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer mysql.CloseConnection(db)

	return router.Run(db)
}
