package logger

import (
	"log"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (l *Mock) Error(err error) {
	log.Println("ERROR: ", err.Error())
}

func (l *Mock) Warn(msg string) {
	log.Println("WARNING: ", msg)
}

func (l *Mock) Info(msg string) {
	log.Println("INFO: ", msg)
}
