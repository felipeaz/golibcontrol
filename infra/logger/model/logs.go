package model

import (
	"time"
)

type Log struct {
	Level   string    `json:"severity"`
	Service string    `json:"service"`
	Message string    `json:"message"`
	Error   string    `json:"error"`
	Time    time.Time `json:"timestamp"`
}
