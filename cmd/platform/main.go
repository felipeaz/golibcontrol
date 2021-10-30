package main

import (
	"log"
	"os"

	"github.com/FelipeAz/golibcontrol/build/server/platform/server"
	_log "github.com/FelipeAz/golibcontrol/infra/logger"
	"github.com/FelipeAz/golibcontrol/infra/mysql/platform/database"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
)

const (
	ServiceName = "Platform Service"
)

func main() {
	db, err := database.Connect(
		os.Getenv("PLATFORM_DB_USER"),
		os.Getenv("PLATFORM_DB_PASSWORD"),
		os.Getenv("PLATFORM_DB_HOST"),
		os.Getenv("PLATFORM_DB_PORT"),
		os.Getenv("PLATFORM_DB_DATABASE"),
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
