package main

import (
	"github.com/FelipeAz/golibcontrol/internal/app/server"
	"log"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Println(err.Error())
	}
}
