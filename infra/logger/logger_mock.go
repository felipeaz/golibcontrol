package logger

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func (l *Mock) Error(err error) {
	fmt.Println("ERROR: ", err.Error())
}

func (l *Mock) Warn(msg string) {
	fmt.Println("WARNING: ", msg)
}

func (l *Mock) Info(msg string) {
	fmt.Println("INFO: ", msg)
}
