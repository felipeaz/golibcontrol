package main

import (
	"github.com/FelipeAz/golibcontrol/build/server/account/server"
	"github.com/FelipeAz/golibcontrol/infra/consumer"
	"github.com/FelipeAz/golibcontrol/infra/http/client"
	"github.com/FelipeAz/golibcontrol/infra/http/request"
	_log "github.com/FelipeAz/golibcontrol/infra/logger"
	"github.com/FelipeAz/golibcontrol/infra/mysql/account/database"
	"github.com/FelipeAz/golibcontrol/infra/mysql/service"
	"github.com/FelipeAz/golibcontrol/infra/redis"
	"log"
	"os"
)

const (
	ServiceName = "Account Service"
	AuthPrefix  = "auth"
)

var (
	envs = map[string]string{
		"ACCOUNT_DB_USER":     "",
		"ACCOUNT_DB_PASSWORD": "",
		"ACCOUNT_DB_HOST":     "",
		"ACCOUNT_DB_PORT":     "",
		"ACCOUNT_DB_DATABASE": "",
		"LOG_FILE":            "",
		"REDIS_HOST":          "",
		"REDIS_PORT":          "",
		"REDIS_EXPIRE":        "",
		"CONSUMERS_HOST":      "",
		"JWT_SECRET_KEY":      "",
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
		envs["ACCOUNT_DB_USER"],
		envs["ACCOUNT_DB_PASSWORD"],
		envs["ACCOUNT_DB_HOST"],
		envs["ACCOUNT_DB_PORT"],
		envs["ACCOUNT_DB_DATABASE"],
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer database.CloseConnection(db)

	logger := _log.NewLogger(envs["LOG_FILE"], ServiceName)
	dbService := service.NewMySQLService(db, logger)

	cache, err := redis.NewCache(
		envs["REDIS_HOST"],
		envs["REDIS_PORT"],
		envs["REDIS_EXPIRE"],
		AuthPrefix,
		logger,
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	requestClient := request.NewHttpRequest(client.NewHTTPClient(), envs["CONSUMERS_HOST"])
	consumersService := consumer.NewConsumer(requestClient, logger, envs["JWT_SECRET_KEY"])

	err = server.Start(dbService, cache, consumersService, logger)
	if err != nil {
		log.Fatal(err.Error())
	}
}
