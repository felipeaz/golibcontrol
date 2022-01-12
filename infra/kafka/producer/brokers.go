package producer

import kafka "github.com/Shopify/sarama"

func (p Producer) Connect() (kafka.SyncProducer, error) {
	return kafka.NewSyncProducer(p.Brokers, p.Config())
}
