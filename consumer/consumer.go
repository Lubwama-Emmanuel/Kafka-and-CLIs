package consumer

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/config"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/models"
)

//go:generate mockgen -destination=mocks/mock_consumer.go -package=mocks . Provider
type Provider interface {
	SetUp(config config.ProviderConfig) error
	Subscribe(topic string) error
	ReadMessage(time.Duration) (models.Message, error)
	Close() error
}

type Consumer struct {
	provider Provider
	run      bool
}

func NewConsumer(provider Provider) *Consumer {
	return &Consumer{provider: provider}
}

func (c *Consumer) SetUpProvider(config config.ProviderConfig) error {
	return c.provider.SetUp(config)
}

func (c *Consumer) ConsumeMessages(topic string) error {
	subErr := c.provider.Subscribe(topic)
	if subErr != nil {
		return fmt.Errorf("subscription error %w", subErr)
	}

	for c.run {
		msg, err := c.provider.ReadMessage(time.Millisecond * 100)
		if err != nil {
			c.provider.Close()
			return fmt.Errorf("read messages error %w", subErr)
		}

		log.Printf("Message: %s\n", string(msg.Value))

		if topic == "test_topic" {
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
