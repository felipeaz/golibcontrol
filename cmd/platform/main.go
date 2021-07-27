package main

import (
	"log"
	"os"

	"github.com/FelipeAz/golibcontrol/build/server/platform/server"
)

func main() {
	err := server.Start(
		os.Getenv("PLATFORM_DB_USER"),
		os.Getenv("PLATFORM_DB_PASSWORD"),
		os.Getenv("PLATFORM_DB_HOST"),
		os.Getenv("PLATFORM_DB_PORT"),
		os.Getenv("PLATFORM_DB_DATABASE"),
	)
	if err != nil {
		log.Println(err.Error())
	}
}
