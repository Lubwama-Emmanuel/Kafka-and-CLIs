package producer

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=mocks/mock_producer.go -package=mocks . Provider

type Provider interface {
	Produce(topic, message string) error
	Flush(timeoutMs int)
	KafkaMessage() error
}

type Messsage struct {
	TopicPartition string
	Value          []byte
}

type Producer struct {
	provider Provider
}

func NewProducer(provider Provider) *Producer {
	return &Producer{provider: provider}
}

func (p *Producer) ProduceMessages(topic, message string) error {
	// Delivery report handler for produced messages.
	go func() {
		err := p.provider.KafkaMessage()
		if err != nil {
			log.Error("an error occurred")
		}
	}()

	// Produce message to topic (asynchronously).
	if err := p.provider.Produce(topic, message); err != nil {
		return fmt.Errorf("an error occurred producing messages %w", err)
	}
	// Wait for message deliveries before shutting down.
	p.provider.Flush(15 * 1000)

	return nil
}
