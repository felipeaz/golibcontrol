package logs

import (
	"time"
)

type Log struct {
	// User string `json:"user"`
	Message string    `json:"message"`
	Time    time.Time `json:"timestamp"`
}
