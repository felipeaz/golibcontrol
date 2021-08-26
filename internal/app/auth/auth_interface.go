package auth

import (
	"github.com/FelipeAz/golibcontrol/infra/auth/model"
)

type AuthInterface interface {
	CreateConsumer(username string) (*model.Consumer, error)
	CreateConsumerKey(consumerId, secret string) (*model.ConsumerKey, error)
	GetConsumerKey(consumerId, keyId string) (*model.ConsumerKey, error)
	GetAllConsumerKeys(consumerId string) (*model.Keys, error)
	RetrieveConsumerKey(consumerId, secret string) (*model.ConsumerKey, error)
	DeleteConsumerKey(consumerId, consumerKeyId string) error
	DeleteConsumer(consumerId string) error
}
