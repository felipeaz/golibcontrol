package producer

import (
	kafka "github.com/Shopify/sarama"
)

type Producer struct {
	Brokers []string
}

type ProducerInterface interface {
	Config() *kafka.Config
	Connect() (kafka.SyncProducer, error)
	Push(topic string, message []byte) error
}
