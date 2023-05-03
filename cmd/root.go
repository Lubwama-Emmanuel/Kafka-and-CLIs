package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumer"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producer"
)

type CMD struct {
	consumer consumer.Consumer
	producer producer.Producer
}

func NewCMD(consumer consumer.Consumer, producer producer.Producer) *CMD {
	return &CMD{
		consumer: consumer,
		producer: producer,
	}
}

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "Kafka-and-CLIs",
	Short: "CLI with Kafka",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
