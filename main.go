package main

// Necessary libs:
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/mysql

import (
	"log"

	"github.com/FelipeAz/golibcontrol/database"
)

func main() {
	db := database.Connect()
	sql, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	defer sql.Close()
}
