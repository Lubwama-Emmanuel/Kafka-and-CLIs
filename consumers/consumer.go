package consumers

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

func Consumer(topic, server, from string) error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "myGroup",
		"auto.offset.reset": from,
	})
	if err != nil {
		return fmt.Errorf("error creating consumer %w", err)
	}

	subErr := consumer.SubscribeTopics([]string{topic}, nil)
	if subErr != nil {
		return fmt.Errorf("error subscribing to topics %w", subErr)
	}

	run := true
	for run {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			consumer.Close()
			return fmt.Errorf("an error occurred %w", err)
		}

		log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	}

	return nil
}
