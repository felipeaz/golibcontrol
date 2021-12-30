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

var (
	envs = map[string]string{
		"MANAGEMENT_DB_USER":     "",
		"MANAGEMENT_DB_PASSWORD": "",
		"MANAGEMENT_DB_HOST":     "",
		"MANAGEMENT_DB_PORT":     "",
		"MANAGEMENT_DB_DATABASE": "",
		"LOG_FILE":               "",
	}
)

func init() {
	for env, _ := range envs {
		var exist bool
		if envs[env], exist = os.LookupEnv(env); !exist {
			log.Fatalf("missing environment variable")
		}
	}
}

func main() {
	db, err := database.Connect(
		envs["MANAGEMENT_DB_USER"],
		envs["MANAGEMENT_DB_PASSWORD"],
		envs["MANAGEMENT_DB_HOST"],
		envs["MANAGEMENT_DB_PORT"],
		envs["MANAGEMENT_DB_DATABASE"],
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.CloseConnection(db)

	logger := _log.NewLogger(envs["LOG_FILE"], ServiceName)
	dbService := service.NewMySQLService(db, logger)

	err = server.Start(dbService, logger)
	if err != nil {
		log.Fatal(err.Error())
	}
}
