package redis

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/FelipeAz/golibcontrol/infra/jwt/model"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/garyburd/redigo/redis"
)

// Cache implements Redis functions
type Cache struct{}

// NewCache returns an instance of Cache
func NewCache() *Cache {
	return &Cache{}
}

// Connect initialize the cache.
func (c *Cache) connect() (redis.Conn, error) {
	h := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	conn, err := redis.Dial("tcp", h)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return conn, nil
}

// Set inserts a value into the Cache
func (c *Cache) Set(key string, value interface{}) error {
	conn, err := c.connect()
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	_, err = conn.Do("SET", key, value)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = conn.Do("EXPIRE", key, os.Getenv("REDIS_EXPIRE"))
	if err != nil {
		log.Println(err)
	}

	return err
}

// Get returns a value from Cache
func (c *Cache) Get(key string) ([]byte, error) {
	conn, err := c.connect()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}

// Flush removes a value from Cache
func (c *Cache) Flush(key string) error {
	conn, err := c.connect()
	if err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	_, err = conn.Do("DEL", key)
	return err
}

// StoreAuth persists the access & refresh token on redis
func (c *Cache) StoreAuth(userid uint, td model.TokenDetails) *errors.ApiError {
	conn, err := c.connect()
	if err != nil {
		log.Println(err)
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.AuthenticationFailMessage,
			Error:   err.Error(),
		}
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	accessDetails := model.AccessDetails{
		AccessUuid:  td.AccessUuid,
		RefreshUuid: td.RefreshUuid,
		UserId:      userid,
	}

	errAccess := c.Set(strconv.Itoa(int(userid)), accessDetails)
	if errAccess != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.AuthenticationFailMessage,
			Error:   errAccess.Error(),
		}
	}

	return nil
}
