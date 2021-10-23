package auth

import (
	"github.com/FelipeAz/golibcontrol/infra/auth/model"
	"github.com/stretchr/testify/mock"
)

type MockAuth struct {
	mock.Mock
}

func (a *MockAuth) CreateConsumer(username string) (*model.Consumer, error) {
	resp := a.Called(username)
	return resp.Get(0).(*model.Consumer), resp.Error(1)
}

func (a *MockAuth) CreateConsumerKey(consumerId, secret string) (*model.ConsumerKey, error) {
	resp := a.Called(consumerId, secret)
	return resp.Get(0).(*model.ConsumerKey), resp.Error(1)
}

func (a *MockAuth) GetConsumerKey(consumerId, keyId string) (*model.ConsumerKey, error) {
	resp := a.Called(consumerId, keyId)
	return resp.Get(0).(*model.ConsumerKey), resp.Error(1)
}

func (a *MockAuth) GetAllConsumerKeys(consumerId string) (*model.Keys, error) {
	resp := a.Called(consumerId)
	return resp.Get(0).(*model.Keys), resp.Error(1)
}

func (a *MockAuth) RetrieveConsumerKey(consumerId, secret string) (*model.ConsumerKey, error) {
	resp := a.Called(consumerId, secret)
	return resp.Get(0).(*model.ConsumerKey), resp.Error(1)
}

func (a *MockAuth) DeleteConsumerKey(consumerId, consumerKeyId string) error {
	resp := a.Called(consumerId, consumerKeyId)
	return resp.Error(0)
}

func (a *MockAuth) DeleteConsumer(consumerId string) error {
	resp := a.Called(consumerId)
	return resp.Error(0)
}
