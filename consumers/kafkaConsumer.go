package consumers

import (
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

func NewKafkaConsumer(server, from string) (*KafkaConsumer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "latest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{consumer}, nil
}

func (c *KafkaConsumer) Subscribe(topic string) error {
	return c.Consumer.SubscribeTopics([]string{topic}, nil)
}

func (c *KafkaConsumer) ReadMessage(time.Duration) (*kafka.Message, error) {
	msg, err := c.Consumer.ReadMessage(-1)
	return msg, err
}

func (c *KafkaConsumer) Close() error {
	return c.Consumer.Close()
}
