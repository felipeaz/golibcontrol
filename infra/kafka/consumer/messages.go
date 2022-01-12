package consumer

import (
	"fmt"
	kafka "github.com/Shopify/sarama"
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
				fmt.Println(err)
			case msg := <-consumer.Messages():
				fmt.Printf("Received message | Topic(%s) | Message(%s) \n", msg.Topic, string(msg.Value))
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
}
