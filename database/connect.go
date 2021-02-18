package database

import (
	"log"

	"github.com/FelipeAz/golibcontrol/database/schema"
	"gorm.io/gorm"
)

// Connect opens a connection to the MYSQL server and set Use to a given table
func Connect() *gorm.DB {
	return schema.Create()
}

// CloseConnection closes the connection to the MYSQL server.
func CloseConnection(db *gorm.DB) {
	sql, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sql.Close()
}
