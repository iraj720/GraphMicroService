package main

import (
	"graph/internal/app/broker_app/broker/cmd/broker"
	"graph/internal/app/destination/cmd/destination"
	"graph/internal/app/reciever/cmd/reciever"
	"graph/internal/app/sender/cmd/sender"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	sender.Register(rootCmd)
	reciever.Register(rootCmd)
	broker.Register(rootCmd)
	destination.Register(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("error while executing your CLI : %v", err)
	}
}
