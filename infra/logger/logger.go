package logger

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/logs"
)

func getLogFile() (f *os.File) {
	filePath, err := filepath.Abs(os.Getenv("LOG_FILE"))
	if err != nil {
		log.Println("failed to retrieve log file:", err.Error())
	}

	f, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		f, err = os.Create(filePath)
		if err != nil {
			log.Println("failed to create log file:", err.Error())
		}
	}

	return
}

func LogError(err error) {
	f := getLogFile()
	defer f.Close()

	errorLog := logs.Log{
		Time:    time.Now(),
		Message: err.Error(),
	}

	b, e := json.Marshal(errorLog)
	if e != nil {
		log.Println(e.Error())
		return
	}

	_, e = f.Write(b)
	if e != nil {
		log.Println(e.Error())
		return
	}
}
