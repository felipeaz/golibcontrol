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
	// filePath, err := filepath.Abs(os.Getenv("LOG_FILE"))
	filePath, err := filepath.Abs("internal/app/logs/log.txt")
	if err != nil {
		log.Println(err.Error())
	}

	f, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		f, err = os.Create(filePath)
		if err != nil {
			log.Println(err.Error())
		}
	}

	return
}

func LogError(err error) {
	f := getLogFile()
	defer f.Close()

	errorLog := logs.Log{
		Message: err.Error(),
		Time:    time.Now(),
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
