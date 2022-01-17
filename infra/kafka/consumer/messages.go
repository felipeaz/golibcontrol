package consumer

import (
	kafka "github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func (c Consumer) Consume(topic string, partition int) error {
	worker, err := c.Connect()
	if err != nil {
		return nil
	}
	consumer, err := worker.ConsumePartition(topic, int32(partition), kafka.OffsetOldest)
	if err != nil {
		return err
	}
	go c.Read(consumer)
	return nil
}

func (c Consumer) Read(consumer kafka.PartitionConsumer) {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				log.Printf("Received message | Topic(%s) | Message(%s) \n", msg.Topic, string(msg.Value))
			case <-sigchan:
				log.Println("Interrupt detected")
				doneCh <- struct{}{}
			}
		}
	}()
}
