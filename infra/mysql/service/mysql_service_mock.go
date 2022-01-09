package service

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockMySQLService struct {
	mock.Mock
}

func (s *MockMySQLService) Find(tx *gorm.DB, domainObj interface{}) (interface{}, error) {
	resp := s.Called(tx, domainObj)
	return resp.Get(0), resp.Error(1)
}

func (s *MockMySQLService) FindOne(tx *gorm.DB, domainObj interface{}) (interface{}, error) {
	resp := s.Called(tx, domainObj)
	return resp.Get(0), resp.Error(1)
}

func (s *MockMySQLService) Persist(domainObj interface{}) error {
	resp := s.Called(domainObj)
	return resp.Error(0)
}

func (s *MockMySQLService) Refresh(tx *gorm.DB, domainObj interface{}) error {
	resp := s.Called(tx, domainObj)
	return resp.Error(0)
}

func (s *MockMySQLService) Set(tx *gorm.DB, domainObj interface{}, field string, value interface{}) error {
	resp := s.Called(tx, domainObj, field, value)
	return resp.Error(0)
}

func (s *MockMySQLService) Remove(domainObj interface{}, id string) *errors.ApiError {
	resp := s.Called(domainObj, id)
	return resp.Get(0).(*errors.ApiError)
}

func (s *MockMySQLService) Preload(preload ...string) *gorm.DB {
	resp := s.Called(preload)
	return resp.Get(0).(*gorm.DB)
}

func (s *MockMySQLService) Join(tx *gorm.DB, join ...string) *gorm.DB {
	resp := s.Called(tx, join)
	return resp.Get(0).(*gorm.DB)
}

func (s *MockMySQLService) Group(tx *gorm.DB, group ...string) *gorm.DB {
	resp := s.Called(tx, group)
	return resp.Get(0).(*gorm.DB)
}

func (s *MockMySQLService) Where(tx *gorm.DB, where string) *gorm.DB {
	resp := s.Called(tx, where)
	return resp.Get(0).(*gorm.DB)
}

func (s *MockMySQLService) GetErrorStatusCode(err error) int {
	resp := s.Called(err)
	return resp.Int(0)
}
