package producer

import (
	kafka "github.com/Shopify/sarama"
)

func (p Producer) Produce(topic string, message []byte) error {
	producer, err := p.Connect()
	if err != nil {
		return err
	}
	defer func(producer kafka.SyncProducer) {
		err := producer.Close()
		if err != nil {

		}
	}(producer)

	msg := &kafka.ProducerMessage{
		Topic: topic,
		Value: kafka.ByteEncoder(message),
	}

	// partition, offset and error
	_, _, err = producer.SendMessage(msg)
	return err
}
