package producers

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

func Producer(topic, server, message string) error {
	producer, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})

	defer producer.Close()

	// Delivery report handler for produced messages.
	go func() {
		for e := range producer.Events() {
			if ev, ok := e.(*kafka.Message); ok {
				if ev.TopicPartition.Error != nil {
					log.Error("Delivery failed:", ev.TopicPartition.Error)
				} else {
					log.Printf("Message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce message to topic (asynchronously).
	if err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil); err != nil {
		return fmt.Errorf("an error occurred producing messages %w", err)
	}
	// Wait for message deliveries before shutting down.
	producer.Flush(15 * 1000)

	return nil
}
