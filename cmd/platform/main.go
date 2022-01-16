package main

import (
	"github.com/FelipeAz/golibcontrol/infra/middleware"
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

var (
	envs = map[string]string{
		"PLATFORM_DB_USER":         "",
		"PLATFORM_DB_PASSWORD":     "",
		"PLATFORM_DB_HOST":         "",
		"PLATFORM_DB_PORT":         "",
		"PLATFORM_DB_DATABASE":     "",
		"PLATFORM_ALLOWED_ORIGINS": "",
		"PLATFORM_ALLOWED_HEADERS": "",
		"LOG_FILE":                 "",
		"KAFKA_BROKER_CONNECT":     "",
	}
)

func init() {
	for env := range envs {
		var exist bool
		if envs[env], exist = os.LookupEnv(env); !exist {
			log.Fatalf("missing environment variable")
		}
	}
}

func main() {
	db, err := database.Connect(
		envs["PLATFORM_DB_USER"],
		envs["PLATFORM_DB_PASSWORD"],
		envs["PLATFORM_DB_HOST"],
		envs["PLATFORM_DB_PORT"],
		envs["PLATFORM_DB_DATABASE"],
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.CloseConnection(db)

	logger := _log.NewLogger(envs["LOG_FILE"], ServiceName)
	dbService := service.NewMySQLService(db, logger)

	mwr := middleware.New(envs["PLATFORM_ALLOWED_ORIGINS"], envs["PLATFORM_ALLOWED_HEADERS"])
	err = server.Start(dbService, logger, mwr)
	if err != nil {
		log.Fatal(err.Error())
	}
}
