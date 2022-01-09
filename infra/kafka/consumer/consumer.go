package consumer

import kafka "github.com/Shopify/sarama"

type Consumer struct {
	Brokers []string
}

func New(brokers []string) *Consumer {
	return &Consumer{
		Brokers: brokers,
	}
}

type ConsumerInterface interface {
	Config() *kafka.Config
	Connect() (kafka.Consumer, error)
	Consume(topic string, partition int) error
}
