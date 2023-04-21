package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "Kafka-and-CLIs",
	Short: "CLI with Kafka",
	Run:   RootCmdRun,
}

func RootCmdRun(cmd *cobra.Command, args []string) {
	log.Info("Hello, World!")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
