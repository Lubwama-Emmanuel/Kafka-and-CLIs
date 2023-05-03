//nolint:errcheck
package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/config"
)

func (c *CMD) SendCmdRun(cmd *cobra.Command, args []string) error {
	channel, _ := cmd.Flags().GetString("channel")
	server, _ := cmd.Flags().GetString("server")
	group, _ := cmd.Flags().GetString("group")
	message, _ := cmd.Flags().GetString("message")

	// creating a producer instance from the received flags
	configs := config.ProviderConfig{
		Server: server,
	}

	err := c.producer.SetUpProvider(configs)
	if err != nil {
		return fmt.Errorf("failed to setup producer provider: %w", err)
	}

	if err := c.producer.ProduceMessages(channel, message); err != nil {
		log.Error("an error occurred: ", err)
	}

	log.Info("You have decided to send to the channel: ", channel)
	log.Info("You are sending through the server: ", server)
	log.Info("You are sending through the group: ", group)
	log.Info("Message sent: ", message)

	return nil
}

func (c *CMD) SendInit() {
	sendCmd := &cobra.Command{
		Use:   "send",
		Short: "Command for producer to send a message",
		Long:  `Command for producer to send a message`,
		RunE:  c.SendCmdRun,
	}
	rootCmd.AddCommand(sendCmd)

	sendflags := []struct {
		flagName string
		desc     string
	}{
		{"channel", "Name of the channel you sending to"},
		{"server", "Port for communication"},
		{"message", "Message to send"},
		{"group", "Group producer is communicating to"},
	}

	for i := range sendflags {
		sendCmd.PersistentFlags().String(sendflags[i].flagName, "", sendflags[i].desc)
		sendCmd.MarkPersistentFlagRequired(sendflags[i].flagName)
	}
}
