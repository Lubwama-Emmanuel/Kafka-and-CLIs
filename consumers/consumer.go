package consumers

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

func Consumer(topic, server, from string) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "myGroup",
		"auto.offset.reset": from,
	})

	if err != nil {
		panic(err)
	}

	if err := c.SubscribeTopics([]string{"development"}, nil); err != nil {
		log.Error("An error occured", err)
	}

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			log.Info("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			errorMessage := fmt.Sprintf("Consumer error: %v (%v)\n", err, msg)
			log.Error(errorMessage)
		}
	}

	c.Close()
}
