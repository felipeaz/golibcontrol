package logger

import (
	"github.com/stretchr/testify/mock"
)

type LoggerMock struct {
	mock.Mock
}

func (l *LoggerMock) Error(err error) {
	l.Called(err)
}
