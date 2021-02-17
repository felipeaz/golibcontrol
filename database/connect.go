package database

/* This package require go-sql-driver and a MYSQL Server installed.
To install the go-sql-driver, please run the following command:
go get -u github.com/go-sql-driver/mysql */

import (
	"github.com/FelipeAz/golibcontrol/database/schema"
	"gorm.io/gorm"
)

// Connect opens a connection to the MYSQL server and set Use to a given table
func Connect() *gorm.DB {
	return schema.Create()
}
