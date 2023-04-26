//nolint:errcheck
package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/config"
)

func (c *CMD) ReceiveCmdRun(cmd *cobra.Command, args []string) error {
	channel, _ := cmd.Flags().GetString("channel")
	server, _ := cmd.Flags().GetString("server")
	from, _ := cmd.Flags().GetString("from")
	group, _ := cmd.Flags().GetString("group")

	configs := config.ProviderConfig{
		Server: server,
		Group:  group,
		From:   from,
	}

	err := c.consumer.SetUpProvider(configs)
	if err != nil {
		return fmt.Errorf("failed to setup consumer provider: %w", err)
	}

	if err := c.consumer.ConsumeMessages(channel); err != nil {
		log.Error("an error occurred: ", err)
	}

	log.Info("You have decided to receive from channel: ", channel)
	log.Info("You are receiving from the: ", from)
	log.Info("You are sending through the server: ", server)
	log.Info("You are sending through the group: ", group)

	return nil
}

func (c *CMD) ReceiveInit() {
	recieveCmd := &cobra.Command{
		Use:   "receive",
		Short: "Command for consumer to receive messages",
		Long:  `Command for consumer to receive messages`,
		RunE:  c.ReceiveCmdRun,
	}

	rootCmd.AddCommand(recieveCmd)

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
		recieveCmd.PersistentFlags().String(recieveflags[i].flagName, "", recieveflags[i].desc)
		recieveCmd.MarkPersistentFlagRequired(recieveflags[i].flagName)
	}
}
