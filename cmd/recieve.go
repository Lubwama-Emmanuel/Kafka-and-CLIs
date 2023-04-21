//nolint:errcheck
package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/blockers"
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumers"
)

var RecieveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Command for consumer to receive messages",
	Long:  `Command for consumer to receive messages`,
	RunE:  ReceiveCmdRun,
}

func ReceiveCmdRun(cmd *cobra.Command, args []string) error {
	channel, _ := cmd.Flags().GetString("channel")
	server, _ := cmd.Flags().GetString("server")
	from, _ := cmd.Flags().GetString("from")
	group, _ := cmd.Flags().GetString("group")

	configs := consumers.ConsumerConfig{
		Server: server,
		Group:  group,
		From:   from,
	}

	consumer, err := blockers.NewKafkaConsumer(configs)
	if err != nil {
		log.Error("an error occurred here: ", err)
	}

	consumerInstance := consumers.NewConsumer(consumer)
	consumerInstance.StartConsumer()
	if err := consumerInstance.ConsumeMessages(channel); err != nil { //nolint:wsl
		log.Error("an error occurred: ", err)
	}

	log.Info("You have decided to receive from channel: ", channel)
	log.Info("You are receiving from the: ", from)
	log.Info("You are sending through the server: ", server)
	log.Info("You are sending through the group: ", group)

	return nil
}

func ReceiveInit() {
	rootCmd.AddCommand(RecieveCmd)

	recieveflags := []struct {
		flagName string
		desc     string
	}{
		{"channel", "Name of the channel you receiving from"},
		{"server", "Port for communication"},
		{"from", "The point to start receiving messages from either start|latest"},
		{"group", "A group to receive messages from"},
	}

	for i := range recieveflags {
		RecieveCmd.PersistentFlags().String(recieveflags[i].flagName, "", recieveflags[i].desc)
		RecieveCmd.MarkPersistentFlagRequired(recieveflags[i].flagName)
	}
}
