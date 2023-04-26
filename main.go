package main

import (
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/brokers"
	cmdProvider "github.com/Lubwama-Emmanuel/Kafka-and-CLIs/cmd"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumer"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producer"
)

func main() {
	provider := brokers.NewKafkaBroker()
	consumerInstance := consumer.NewConsumer(provider)
	consumerInstance.StartConsumer()

	producerInstance := producer.NewProducer(provider)

	cmd := cmdProvider.NewCMD(*consumerInstance, *producerInstance)

	cmd.ReceiveInit()
	cmd.SendInit()
	cmdProvider.Execute()
}
