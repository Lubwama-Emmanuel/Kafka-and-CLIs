package main

import "github.com/Lubwama-Emmanuel/Kafka-and-CLIs/cmd"

func main() {
	cmd.RootInit()
	cmd.ReceiveInit()
	cmd.SendInit()
	cmd.Execute()
}
