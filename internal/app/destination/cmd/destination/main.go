package destination

import "github.com/spf13/cobra"

func main() {

}

func Register(cmd *cobra.Command) {
	var sender = &cobra.Command{
		Use:   "startSender",
		Short: "sender client starts to send requests rapidly",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
	cmd.AddCommand(sender)
}
