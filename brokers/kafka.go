package brokers

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/config"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/models"
)

// Kafka struct implementing the consumer interface.
type KafkaBroker struct {
	consumer *kafka.Consumer
	producer *kafka.Producer
}

// Creates new kafka consumer instance.
func NewKafkaBroker() *KafkaBroker {
	return &KafkaBroker{}
}

func (c *KafkaBroker) SetUp(config config.ProviderConfig) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Server,
		"group.id":          config.Group,
		"auto.offset.reset": "latest",
	})
	if err != nil {
		return fmt.Errorf("failed to set up kafka consumer: %w", err)
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Server,
	})
	if err != nil {
		return fmt.Errorf("failed to set up kafka producer: %w", err)
	}

	c.consumer = consumer
	c.producer = producer

	return nil
}

// Subscribes to a given topic.
func (c *KafkaBroker) Subscribe(topic string) error {
	err := c.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	return nil
}

// Reads received messages.
func (c *KafkaBroker) ReadMessage(time.Duration) (models.Message, error) {
	msg, err := c.consumer.ReadMessage(-1)
	if err != nil {
		return models.Message{}, fmt.Errorf("an error occurred: %w", err)
	}

	return models.Message{Value: msg.Value}, nil
}

// Closes the consumer.
func (c *KafkaBroker) Close() error {
	if err := c.consumer.Close(); err != nil {
		return fmt.Errorf("an error occurred: %w", err)
	}

	return nil
}

// produces messages to a given topic.
func (k *KafkaBroker) Produce(topic, message string) error {
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
func (k *KafkaBroker) KafkaMessage() error {
	for e := range k.producer.Events() {
		if ev, ok := e.(*kafka.Message); ok {
			if ev.TopicPartition.Error != nil {
				return fmt.Errorf("delivery failed: %w", ev.TopicPartition.Error)
			}
		}
	}

	return nil
}

// Stops the produces after a given time.
func (k *KafkaBroker) Flush(timeoutMs int) {
	k.producer.Flush(timeoutMs)
}
