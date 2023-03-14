// nolint
package cmd

import (
	"github.com/Lubwama-Emmanuel/Kafka-and-CLIs/consumers"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// recieveCmd represents the recieve command
var recieveCmd = &cobra.Command{
	Use:   "recieve",
	Short: "Command for consumer to recieve messages",
	Long:  `Command for consumer to recieve messages`,

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

func init() {
	rootCmd.AddCommand(recieveCmd)

	recieveCmd.PersistentFlags().String("channel", "", "Name of the channel you receiving from")
	recieveCmd.MarkPersistentFlagRequired("channel")
	recieveCmd.PersistentFlags().String("server", "", "Port for communication")
	recieveCmd.MarkPersistentFlagRequired("server")
	recieveCmd.PersistentFlags().String("from", "", "The point to start receiving messages from either start|latest")
	recieveCmd.MarkPersistentFlagRequired("from")
	recieveCmd.PersistentFlags().String("group", "", "A group to receive messages from. Group is optional")

}
