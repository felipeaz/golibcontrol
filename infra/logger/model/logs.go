package model

import (
	"time"
)

type Log struct {
	Message string    `json:"message"`
	Service string    `json:"service"`
	Error   string    `json:"error"`
	Level   string    `json:"severity"`
	Time    time.Time `json:"timestamp"`
}
