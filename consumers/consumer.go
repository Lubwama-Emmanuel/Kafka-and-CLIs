package consumers

import (
	"fmt"
	"time"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/models"
	log "github.com/sirupsen/logrus"
)

type Provider interface {
	SetUp(ConsumerConfig) error
	SubscribeTopics([]string) error
	ReadMessage(time.Duration) (models.Message, error)
	Close() error
}

type ConsumerConfig struct {
	Host   string
	Group  string
	Offset string
}

type Consumer struct {
	provider Provider
	run      bool
}

func NewConsumer(provider Provider) *Consumer {
	return &Consumer{provider: provider}
}

func (c *Consumer) Consume(topic, from string) error {
	subErr := c.provider.SubscribeTopics([]string{topic})
	if subErr != nil {
		return fmt.Errorf("error subscribing to topics %w", subErr)
	}

	for c.run {
		msg, readErr := c.provider.ReadMessage(time.Millisecond * 100)
		if readErr != nil {
			closeErr := c.provider.Close()
			if closeErr != nil {
				return fmt.Errorf("error closing consumer %w", closeErr)
			}

			return fmt.Errorf("an error occurred while reading from provider %w", readErr)
		}

		log.Printf("Message on %s: %s\n", msg.Topic, string(msg.Value))

		if topic == "test-topic" {
			c.StopConsumer()
		}
	}

	return nil
}

func (c *Consumer) StartConsumer() {
	c.run = true
}

func (c *Consumer) StopConsumer() {
	c.run = false
}
