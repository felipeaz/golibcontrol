package consumer

import kafka "github.com/Shopify/sarama"

func (c Consumer) Connect() (kafka.Consumer, error) {
	return kafka.NewConsumer(c.Brokers, c.Config())
}
