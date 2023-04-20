package blockers

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumers"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/models"
)

// Kafka struct implementing the consumer interface.
type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

// Creates new kafka consumer instance.
func NewKafkaConsumer(config consumers.ConsumerConfig) (*KafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Server,
		"group.id":          config.Group,
		"auto.offset.reset": config.From,
	})
	if err != nil {
		return nil, fmt.Errorf("an error occurred: %w", err)
	}

	return &KafkaConsumer{Consumer: consumer}, nil
}

// Subscribes to a given topic.
func (c *KafkaConsumer) Subscribe(topic string) error {
	err := c.Consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	return nil
}

// Reads received messages.
func (c *KafkaConsumer) ReadMessage(time.Duration) (models.Message, error) {
	msg, err := c.Consumer.ReadMessage(-1)
	if err != nil {
		return models.Message{}, fmt.Errorf("an error occurred: %w", err)
	}

	return models.Message{Value: msg.Value}, nil
}

// Closes the consumer.
func (c *KafkaConsumer) Close() error {
	if err := c.Consumer.Close(); err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	return nil
}

// Kafka struct that implements the producer interface.
type KafkaProducer struct {
	producer *kafka.Producer
}

// Creates new kafka producer instance.
func NewKafkaProducer(server string) (*KafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": server,
	})
	if err != nil {
		return nil, fmt.Errorf("an error occurred: %w", err)
	}

	return &KafkaProducer{producer}, nil
}

// produces messages to a given topic.
func (k *KafkaProducer) Produce(topic, message string) error {
	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	if err != nil {
		return fmt.Errorf("produce err %w", err)
	}

	return nil
}

// creates produce events.
func (k *KafkaProducer) Events() chan kafka.Event {
	return k.producer.Events()
}

// Stops the produces after a given time.
func (k *KafkaProducer) Flush(timeoutMs int) {
	k.producer.Flush(timeoutMs)
}
