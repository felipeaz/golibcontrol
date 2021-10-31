package logger

import (
	"fmt"

	"github.com/stretchr/testify/mock"
)

type LoggerMock struct {
	mock.Mock
}

func (l *LoggerMock) Error(err error) {
	fmt.Println("log error: ", err.Error())
}
