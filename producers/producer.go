package producers

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

func Producer(topic, server, message string) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		log.Panic(err)
	}

	defer producer.Close()

	// Delivery report handler for produced messages.
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Delivery failed: %v\n", ev.TopicPartition)
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
		log.Error("An error occurred", err)
	}

	// Wait for message deliveries before shutting down.
	producer.Flush(15 * 1000)
}
