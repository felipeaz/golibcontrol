package main

import (
	"log"
	"os"

	"github.com/FelipeAz/golibcontrol/build/server/management/server"
	_log "github.com/FelipeAz/golibcontrol/infra/logger"
	"github.com/FelipeAz/golibcontrol/infra/mysql/management/database"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
)

const (
	ServiceName = "Management Service"
)

func main() {
	db, err := database.Connect(
		os.Getenv("MANAGEMENT_DB_USER"),
		os.Getenv("MANAGEMENT_DB_PASSWORD"),
		os.Getenv("MANAGEMENT_DB_HOST"),
		os.Getenv("MANAGEMENT_DB_PORT"),
		os.Getenv("MANAGEMENT_DB_DATABASE"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.CloseConnection(db)

	logger := _log.NewLogger(os.Getenv("LOG_FILE"), ServiceName)
	dbService := service.NewMySQLService(db, logger)

	err = server.Start(dbService, logger)
	if err != nil {
		log.Fatal(err.Error())
	}
}
