package consumers

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer interface {
	NewKafkaConsumer(server, topic string) (*KafkaConsumer, error)
	Subscribe(topic string) error
	ReadMessage(time.Duration) (*kafka.Message, error)
	Close() error
}

type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

func NewKafkaConsumer(server, group, from string) (*KafkaConsumer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          group,
		"auto.offset.reset": from,
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, fmt.Errorf("an error occurred: %w", err)
	}

	return &KafkaConsumer{consumer}, nil
}

func (c *KafkaConsumer) Subscribe(topic string) error {
	err := c.Consumer.SubscribeTopics([]string{topic}, nil)
	return fmt.Errorf("an error occurred: %w", err)
}

func (c *KafkaConsumer) ReadMessage(time.Duration) (*kafka.Message, error) {
	msg, err := c.Consumer.ReadMessage(-1)
	if err != nil {
		return nil, fmt.Errorf("an error occurred: %w", err)
	}

	return msg, nil
}

func (c *KafkaConsumer) Close() error {
	if err := c.Consumer.Close(); err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	return nil
}

func StartConsumer() bool {
	return true
}

func StopConsumer() bool {
	return false
}
