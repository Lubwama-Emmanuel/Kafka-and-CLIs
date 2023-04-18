package producers

import (
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
		return nil, err
	}

	return &KafkaProducer{producer}, err
}

func (p *KafkaProducer) Close() {
	p.producer.Close()
}

func (p *KafkaProducer) Events() chan kafka.Event {
	return p.producer.Events()
}

func (p *KafkaProducer) Send(topic, message string) error {
	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
}

func (p *KafkaProducer) Flush(timeoutMs int) int {
	return p.producer.Flush(timeoutMs)
}
