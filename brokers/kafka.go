package consumers

import (
	"fmt"
	"time"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumers"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/models"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KakfaProvider struct {
	config   consumers.ConsumerConfig
	consumer *kafka.Consumer
}

func (k *KakfaProvider) SetUp(config consumers.ConsumerConfig) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": k.config.Host,
		"group.id":          k.config.Group,
		"auto.offset.reset": k.config.Offset,
	})
	if err != nil {
		return fmt.Errorf("error creating consumer %w", err)
	}

	k.consumer = consumer

	return nil
}

func (k *KakfaProvider) SubscribeTopics(topics []string) error {
	return k.consumer.SubscribeTopics(topics, nil)
}

func (k *KakfaProvider) ReadMessage(timeout time.Duration) (models.Message, error) {
	msg, readErr := k.consumer.ReadMessage(-1)
	if readErr != nil {
		return models.Message{}, fmt.Errorf("an error occurred while reading from kafka %w", readErr)
	}

	return models.Message{Topic: msg.TopicPartition.String(), Value: msg.Value}, nil
}

func (k *KakfaProvider) Close() error {
	return k.consumer.Close()
}
