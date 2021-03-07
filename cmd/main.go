package main

import (
	"github.com/FelipeAz/golibcontrol/build/server"
	"log"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Println(err.Error())
	}
}
