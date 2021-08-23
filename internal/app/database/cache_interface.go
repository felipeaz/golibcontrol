package database

type CacheInterface interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Flush(key string) error
}
