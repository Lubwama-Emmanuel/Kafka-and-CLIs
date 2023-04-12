//nolint:errcheck
package cmd

import (
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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

	log.Info("You have decided to receive from channel: ", channel)
	log.Info("You are receiving from the: ", from)
	log.Info("You are sending through the server: ", server)
	log.Info("You are sending through the group: ", group)
	consumers.Consumer(channel, server, from)

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
	}

	for i := range recieveflags {
		RecieveCmd.PersistentFlags().String(recieveflags[i].flagName, "", recieveflags[i].desc)
		RecieveCmd.MarkPersistentFlagRequired(recieveflags[i].flagName)
	}

	RecieveCmd.PersistentFlags().String("group", "", "A group to receive messages from. Group is optional")
}
