package server

import (
	"github.com/FelipeAz/golibcontrol/platform/mysql"
	"github.com/FelipeAz/golibcontrol/platform/router"
	"log"
)

// Start initialize the webservice
func Start() (err error) {
	db, err := mysql.Connect()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer mysql.CloseConnection(db)

	err = router.Run()
	return
}
