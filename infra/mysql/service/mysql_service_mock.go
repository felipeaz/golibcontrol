package service

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/stretchr/testify/mock"
)

type MockMySQLService struct {
	mock.Mock
}

func (s *MockMySQLService) FetchAll(domainObj interface{}) (interface{}, *errors.ApiError) {
	resp := s.Called(domainObj)
	return resp.Get(0), resp.Get(1).(*errors.ApiError)
}

func (s *MockMySQLService) FetchAllWithPreload(domainObj interface{}, preload string) (interface{}, *errors.ApiError) {
	resp := s.Called(domainObj, preload)
	return resp.Get(0), resp.Get(1).(*errors.ApiError)
}

func (s *MockMySQLService) Fetch(domainObj interface{}, id string) (interface{}, *errors.ApiError) {
	resp := s.Called(domainObj, id)
	return resp.Get(0), resp.Get(1).(*errors.ApiError)
}

func (s *MockMySQLService) FetchWithPreload(domainObj interface{}, id, preload string) (interface{}, *errors.ApiError) {
	resp := s.Called(domainObj, id, preload)
	return resp.Get(0), resp.Get(1).(*errors.ApiError)
}

func (s *MockMySQLService) FetchAllWhere(domainObj interface{}, fieldName, fieldValue string) (interface{}, *errors.ApiError) {
	resp := s.Called(domainObj, fieldName, fieldValue)
	return resp.Get(0), resp.Get(1).(*errors.ApiError)
}

func (s *MockMySQLService) FetchAllWhereWithQuery(domainObj interface{}, query string) (interface{}, *errors.ApiError) {
	resp := s.Called(domainObj, query)
	return resp.Get(0), resp.Get(1).(*errors.ApiError)
}

func (s *MockMySQLService) Persist(domainObj interface{}) *errors.ApiError {
	resp := s.Called(domainObj)
	return resp.Get(0).(*errors.ApiError)
}

func (s *MockMySQLService) Refresh(domainObj interface{}, id string) *errors.ApiError {
	resp := s.Called(domainObj, id)
	return resp.Get(0).(*errors.ApiError)
}

func (s *MockMySQLService) Remove(domainObj interface{}, id string) *errors.ApiError {
	resp := s.Called(domainObj, id)
	return resp.Get(0).(*errors.ApiError)
}
