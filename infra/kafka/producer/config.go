package producer

import kafka "github.com/Shopify/sarama"

func (p Producer) Config() *kafka.Config {
	cfg := kafka.NewConfig()
	cfg.Producer.Return.Successes = true
	cfg.Producer.RequiredAcks = kafka.WaitForAll
	cfg.Producer.Retry.Max = 5
	return cfg
}
