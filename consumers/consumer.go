package consumers

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	log "github.com/sirupsen/logrus"
)

func Consumer(topic, server, from string) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          "myGroup",
		"auto.offset.reset": from,
	})
	if err != nil {
		panic(err)
	}

	if err := consumer.SubscribeTopics([]string{"development"}, nil); err != nil {
		log.Error("An error occurred", err)
	}

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true

	for run {
		msg, err := consumer.ReadMessage(time.Second)
		if err == nil {
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else if !err.(kafka.Error).IsTimeout() {
			errorMessage := fmt.Sprintf("Consumer error: %v (%v)\n", err, msg)
			log.Error(errorMessage)
		}
	}

	consumer.Close()
}
