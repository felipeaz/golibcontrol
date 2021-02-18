package main

import (
	"github.com/FelipeAz/golibcontrol/database"
)

func main() {
	db := database.Connect()
	defer database.CloseConnection(db)
}
