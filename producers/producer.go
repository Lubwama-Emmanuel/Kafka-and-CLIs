package producers

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=mocks/mock_producer.go -package=pmocks . Provider

type Provider interface {
	Produce(topic, message string) error
	Events() chan kafka.Event
	Flush(timeoutMs int)
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
		for e := range p.provider.Events() {
			if ev, ok := e.(*kafka.Message); ok {
				if ev.TopicPartition.Error != nil {
					log.Error("Delivery failed:", ev.TopicPartition.Error)
				} else {
					log.Printf("Message to %v\n", ev.TopicPartition)
				}
			}
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
