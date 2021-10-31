package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/FelipeAz/golibcontrol/infra/logger/model"
)

const (
	ClosingLogFileErrorMessage = "failed to close log file"
	ErrorLevel                 = "ERROR"
	WarnLevel                  = "WARNING"
	InfoLevel                  = "INFO"
)

type Logger struct {
	logFilePath string
	serviceName string
}

func NewLogger(logFilePath, service string) Logger {
	return Logger{
		logFilePath: logFilePath,
		serviceName: service,
	}
}

func (l Logger) Error(err error) {
	errorLog := model.Log{
		Level:   ErrorLevel,
		Service: l.serviceName,
		Error:   err.Error(),
		Time:    time.Now(),
	}
	l.writeError(errorLog)
}

func (l Logger) Warn(msg string) {
	errorLog := model.Log{
		Level:   WarnLevel,
		Service: l.serviceName,
		Message: msg,
		Time:    time.Now(),
	}
	l.writeError(errorLog)
}

func (l Logger) Info(msg string) {
	errorLog := model.Log{
		Level:   InfoLevel,
		Service: l.serviceName,
		Message: msg,
		Time:    time.Now(),
	}
	l.writeError(errorLog)
}

func (l Logger) getLogFile(path string) (f *os.File) {
	filePath, err := filepath.Abs(path)
	if err != nil {
		log.Println("failed to retrieve log file:", l.serviceName, err.Error())
		return
	}

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		err = os.Mkdir(filePath, 0755)
		if err != nil {
			log.Println("failed to create logs folder", l.serviceName, err.Error())
			return
		}
	}

	fileName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))
	fullPath := fmt.Sprintf("%s/%s", filePath, fileName)

	f, err = os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("failed to create log file:", l.serviceName, err.Error())
		return nil
	}

	return
}

func (l Logger) writeError(errorLog model.Log) {
	f := l.getLogFile(l.logFilePath)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(l.serviceName, ClosingLogFileErrorMessage, err.Error())
		}
	}(f)

	b, e := json.Marshal(errorLog)
	if e != nil {
		log.Println(l.serviceName, e.Error())
	}

	_, e = f.Write(b)
	if e != nil {
		log.Println(l.serviceName, e.Error())
	}
}
