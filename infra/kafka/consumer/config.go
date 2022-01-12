package consumer

import kafka "github.com/Shopify/sarama"

func (c Consumer) Config() *kafka.Config {
	cfg := kafka.NewConfig()
	cfg.Consumer.Return.Errors = true
	return cfg
}
