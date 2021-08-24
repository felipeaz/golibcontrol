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

func getLogFile(path string) (f *os.File) {
	filePath, err := filepath.Abs(path)
	if err != nil {
		log.Println("failed to retrieve log file:", err.Error())
		return
	}

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		err = os.Mkdir(filePath, 0755)
		if err != nil {
			log.Println("failed to create logs folder", err.Error())
			return
		}
	}

	fileName := fmt.Sprintf("%s.log", time.Now().Format("2006-01-02"))
	fullPath := fmt.Sprintf("%s/%s", filePath, fileName)

	f, err = os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("failed to create log file:", err.Error())
		return nil
	}

	return
}

func LogError(err error) {
	f := getLogFile(os.Getenv("LOG_FILE"))
	defer f.Close()

	errorLog := model.Log{
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
