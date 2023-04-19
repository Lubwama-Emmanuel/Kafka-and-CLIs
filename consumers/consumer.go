package consumers

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type Provider interface {
	SetUp(ConsumerConfig) error
	SubscribeTopics([]string) error
	ReadMessage(time.Duration) (Message, error)
	Close() error
}

type ConsumerConfig struct {
	Host   string
	Group  string
	Offset string
}

type Message struct {
	Topic string
	Value []byte
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
		msg, readErr := c.provider.ReadMessage(-1)
		if readErr != nil {
			closeErr := c.provider.Close()
			if closeErr != nil {
				return fmt.Errorf("error closing consumer %w", closeErr)
			}

			return fmt.Errorf("an error occurred while reading from provider %w", readErr)
		}

		log.Printf("Message on %s: %s\n", msg.Topic, string(msg.Value))
	}

	return nil
}

func (c *Consumer) StartConsumer() {
	c.run = true
}

func (c *Consumer) StopConsumer() {
	c.run = false
}
