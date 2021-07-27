package main

import (
	"log"
	"os"

	"github.com/FelipeAz/golibcontrol/build/server/management/server"
)

func main() {
	err := server.Start(
		os.Getenv("MANAGEMENT_DB_USER"),
		os.Getenv("MANAGEMENT_DB_PASSWORD"),
		os.Getenv("MANAGEMENT_DB_HOST"),
		os.Getenv("MANAGEMENT_DB_PORT"),
		os.Getenv("MANAGEMENT_DB_DATABASE"),
	)
	if err != nil {
		log.Println(err.Error())
	}
}
