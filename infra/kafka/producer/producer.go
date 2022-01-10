package producer

import (
	kafka "github.com/Shopify/sarama"
)

type Producer struct {
	Brokers []string
}

func New(brokers ...string) *Producer {
	return &Producer{
		Brokers: brokers,
	}
}

type ProducerInterface interface {
	Config() *kafka.Config
	Connect() (kafka.SyncProducer, error)
	Produce(topic string, message []byte) error
}
