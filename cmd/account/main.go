package main

import (
	"log"
	"os"

	"github.com/FelipeAz/golibcontrol/build/server/account/server"
)

func main() {
	err := server.Start(
		os.Getenv("ACCOUNT_DB_USER"),
		os.Getenv("ACCOUNT_DB_PASSWORD"),
		os.Getenv("ACCOUNT_DB_HOST"),
		os.Getenv("ACCOUNT_DB_PORT"),
		os.Getenv("ACCOUNT_DB_DATABASE"),
		os.Getenv("CONSUMERS_HOST"),
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_EXPIRE"),
	)
	if err != nil {
		log.Println(err.Error())
	}
}
