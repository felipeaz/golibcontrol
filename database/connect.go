package database

import (
	"github.com/FelipeAz/golibcontrol/database/schema"
	"gorm.io/gorm"
)

// Connect opens a connection to the MYSQL server and set Use to a given table
func Connect() *gorm.DB {
	return schema.Create()
}
