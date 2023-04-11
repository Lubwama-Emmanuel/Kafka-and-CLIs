//nolint:errcheck
package cmd

import (
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// receiveCmd represents the receive command.
var recieveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Command for consumer to receive messages",
	Long:  `Command for consumer to receive messages`,

	Run: func(cmd *cobra.Command, args []string) {
		channel, _ := cmd.Flags().GetString("channel")
		server, _ := cmd.Flags().GetString("server")
		from, _ := cmd.Flags().GetString("from")
		group, _ := cmd.Flags().GetString("group")
		log.Info("You have decided to send to the channel: ", channel)
		log.Info("You are receiving from the: ", from)
		log.Info("You are sending through the server: ", server)
		log.Info("You are sending through the group: ", group)
		consumers.Consumer(channel, server, from)
	},
}

func ReceiveInit() {
	rootCmd.AddCommand(recieveCmd)

	recieveflags := []struct {
		flagName string
		desc     string
	}{
		{"channel", "Name of the channel you receiving from"},
		{"server", "Port for communication"},
		{"from", "The point to start receiving messages from either start|latest"},
	}

	for i := range recieveflags {
		recieveCmd.PersistentFlags().String(recieveflags[i].flagName, "", recieveflags[i].desc)
		recieveCmd.MarkPersistentFlagRequired(recieveflags[i].flagName)
	}

	recieveCmd.PersistentFlags().String("group", "", "A group to receive messages from. Group is optional")
}
