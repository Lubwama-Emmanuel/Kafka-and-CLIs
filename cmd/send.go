//nolint:errcheck
package cmd

import (
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/producers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command.
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Command for producer to send a message",
	Long:  `Command for producer to send a message`,
	RunE:  SendCmdRun,
}

func SendCmdRun(cmd *cobra.Command, args []string) error {
	channel, _ := cmd.Flags().GetString("channel")
	server, _ := cmd.Flags().GetString("server")
	group, _ := cmd.Flags().GetString("group")
	message, _ := cmd.Flags().GetString("message")
	producers.Producer(channel, server, message)
	log.Info("You have decided to send to the channel: ", channel)
	log.Info("You are sending through the server: ", server)
	log.Info("You are sending through the group: ", group)
	log.Info("Message sent: ", message)
	return nil
}

func SendInit() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.PersistentFlags().String("channel", "", "Name of the channel you sending to")
	sendCmd.MarkPersistentFlagRequired("channel")
	sendCmd.PersistentFlags().String("server", "", "Port for communication")
	sendCmd.MarkPersistentFlagRequired("server")
	sendCmd.PersistentFlags().String("message", "", "Message to send")
	sendCmd.MarkPersistentFlagRequired("message")
	sendCmd.PersistentFlags().String("group", "", "Gruop producer is communicating to")
}
