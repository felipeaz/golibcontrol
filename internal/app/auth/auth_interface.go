package auth

import (
	"github.com/FelipeAz/golibcontrol/infra/auth/model"
)

type AuthInterface interface {
	GetConsumer(consumerId string) (string, error)
	CreateConsumer(username string) (*model.Consumer, error)
	CreateConsumerKey(consumerId, secret string) (*model.ConsumerKey, error)
	GetConsumerKey(consumerId, keyId string) (*model.ConsumerKey, error)
	GetAllConsumerKeys(consumerId string) (*model.Keys, error)
	RetrieveConsumerKey(consumerId, secret string) (*model.ConsumerKey, error)
	DeleteConsumer(consumerId, consumerKeyId string) error
}
