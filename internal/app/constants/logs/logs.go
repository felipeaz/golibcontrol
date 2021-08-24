package logs

import (
	"time"
)

type Log struct {
	// User string `json:"user"`
	Message string    `json:"message"`
	Service string    `json:"service"`
	Error   string    `json:"error"`
	Time    time.Time `json:"timestamp"`
}
