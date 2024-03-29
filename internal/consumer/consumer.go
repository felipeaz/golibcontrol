package consumer

import (
	"github.com/FelipeAz/golibcontrol/infra/consumer/model"
)

type Interface interface {
	CreateConsumer(username string) (*model.Consumer, error)
	CreateConsumerKey(consumerId string) (*model.ConsumerKey, error)
	GetConsumerKey(consumerId, keyId string) (*model.ConsumerKey, error)
	GetAllConsumerKeys(consumerId string) (*model.Keys, error)
	RetrieveConsumerKey(consumerId string) (*model.ConsumerKey, error)
	DeleteConsumerKey(consumerId, consumerKeyId string) error
	DeleteConsumer(consumerId string) error
}
