package main

import (
	"graph/internal/app/sender/cmd/sender"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}
	sender.Register(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatalf("error while executing your CLI : %v", err)
	}
}
