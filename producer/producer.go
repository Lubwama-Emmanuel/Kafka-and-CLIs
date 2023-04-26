package producer

import (
	"fmt"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/config"
	log "github.com/sirupsen/logrus"
)

//go:generate mockgen -destination=mocks/mock_producer.go -package=mocks . Provider

type Provider interface {
	SetUp(config config.ProviderConfig) error
	Produce(topic, message string) error
	Flush(timeoutMs int)
	DeliveryReport() error
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

func (p *Producer) SetUpProvider(config config.ProviderConfig) error {
	err := p.provider.SetUp(config)
	if err != nil {
		return fmt.Errorf("failed to set up provider %w", err)
	}

	return nil
}

func (p *Producer) ProduceMessages(topic, message string) error {
	// Delivery report handler for produced messages.
	go func() {
		err := p.provider.DeliveryReport()
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
