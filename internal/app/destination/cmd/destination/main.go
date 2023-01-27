package destination

import (
	"graph/internal/app/destination/controller"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9990"
	SERVER_TYPE = "tcp"
)

func main() {
	dc := controller.NewDestinationController()
	err := dc.StartServing(SERVER_HOST, SERVER_PORT, SERVER_TYPE)
	if err != nil {
		logrus.Fatalf("cannot start serving %v", err)
	}
}

func Register(cmd *cobra.Command) {
	var dest = &cobra.Command{
		Use:     "destination",
		Aliases: []string{"dest"},
		Short:   "sender client starts to send requests rapidly",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
	cmd.AddCommand(dest)
}
