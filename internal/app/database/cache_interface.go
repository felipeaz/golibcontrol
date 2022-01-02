package database

type Cache interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Flush(key string) error
}
