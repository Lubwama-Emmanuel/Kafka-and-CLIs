package models

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type Message struct {
	Value []byte
}

type Kmessage struct {
	TopicPartition kafka.TopicPartition
}
