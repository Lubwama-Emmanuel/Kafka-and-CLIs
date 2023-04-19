package producers

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Producer interface {
	Events() chan kafka.Event
	Send(topic, message string) error
	Flush(timeoutMs int) int
}

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(server string) (*KafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		return nil, fmt.Errorf("an error occurred: %w", err)
	}

	return &KafkaProducer{producer}, nil
}

func (p *KafkaProducer) Close() {
	p.producer.Close()
}

func (p *KafkaProducer) Events() chan kafka.Event {
	return p.producer.Events()
}

func (p *KafkaProducer) Send(topic, message string) error {
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	if err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	return nil
}

func (p *KafkaProducer) Flush(timeoutMs int) int {
	return p.producer.Flush(timeoutMs)
}
