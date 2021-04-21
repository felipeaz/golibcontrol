package redis

import (
	"fmt"
	"log"
	"os"

	"github.com/FelipeAz/golibcontrol/platform/logger"
	"github.com/garyburd/redigo/redis"
)

type Cache struct {
	Conn redis.Conn
}

// Connect initialize the cache.
func Connect() (*Cache, error) {
	h := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	conn, err := redis.Dial("tcp", h)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &Cache{
		Conn: conn,
	}, nil
}

// CloseConnection closes the connection with the cache and log an error to the log file if needed
func (c Cache) CloseConnection() {
	err := c.Conn.Close()
	if err != nil {
		logger.LogError(err)
	}
}
