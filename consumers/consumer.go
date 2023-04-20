package consumers

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/models"
)

type Provider interface {
	Subscribe(topic string) error
	ReadMessage(time.Duration) (models.Message, error)
	Close() error
}

type Consumer struct {
	provider Provider
	run      bool
}

type ConsumerConfig struct {
	Server string
	From   string
	Group  string
}

func NewConsumer(provider Provider) *Consumer {
	return &Consumer{provider: provider}
}

func (c *Consumer) ConsumeMessages(topic string) {
	subErr := c.provider.Subscribe(topic)
	if subErr != nil {
		log.Error("subscription error", subErr)
	}

	for c.run {
		msg, err := c.provider.ReadMessage(-1)
		if err != nil {
			c.provider.Close()
			log.Error("read messages error", err)
		}

		log.Printf("Message: %s\n", string(msg.Value))
	}
}

func (c *Consumer) StartConsumer() {
	c.run = true
}

func (c *Consumer) StopConsumer() {
	c.run = false
}
