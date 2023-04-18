package producers

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

func (p *KafkaProducer) ProduceMessages(topic, message string) {
	// Delivery report handler for produced messages.
	go func() {
		for e := range p.producer.Events() {
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
	if err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil); err != nil {
		log.Error("an error occurred producing messages ", err)
	}
	// Wait for message deliveries before shutting down.
	p.producer.Flush(15 * 1000)
}
