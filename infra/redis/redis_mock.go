package redis

import (
	"github.com/stretchr/testify/mock"
)

// MockCache implements Redis functions
type MockCache struct {
	mock.Mock
}

// Set inserts a value into the MockCache
func (c *MockCache) Set(key string, value []byte) error {
	resp := c.Called(key, value)
	return resp.Error(0)
}

// Get returns a value from MockCache
func (c *MockCache) Get(key string) ([]byte, error) {
	resp := c.Called(key)
	return resp.Get(0).([]byte), resp.Error(1)
}

// Flush removes a value from MockCache
func (c *MockCache) Flush(key string) error {
	resp := c.Called(key)
	return resp.Error(0)
}
