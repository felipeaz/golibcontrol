package auth

import (
	"github.com/FelipeAz/golibcontrol/infra/auth/model"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

type AuthInterface interface {
	GetConsumer(consumerId string) (string, *errors.ApiError)
	CreateConsumer(username string) (*model.Consumer, *errors.ApiError)
	CreateConsumerKey(consumerId, secret string) (*model.ConsumerKey, *errors.ApiError)
	GetConsumerKey(consumerId, keyId string) (*model.ConsumerKey, *errors.ApiError)
	GetAllConsumerKeys(consumerId string) (*model.Keys, *errors.ApiError)
	RetrieveConsumerKey(consumerId, secret string) (*model.ConsumerKey, *errors.ApiError)
	DeleteConsumer(consumerId, consumerKeyId string) *errors.ApiError
}
