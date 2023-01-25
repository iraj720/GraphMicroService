package broker

import "github.com/spf13/cobra"

func RegisterBroker(root *cobra.Command) {
	root.AddCommand()
}