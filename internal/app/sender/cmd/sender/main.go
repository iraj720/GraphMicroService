package sender

import (
	"fmt"
	"graph/internal/app/sender/client"
	"graph/internal/app/sender/service"
	"graph/pkg/data_handler"
	"graph/proto/reciever"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := reciever.NewRecieverClient(conn)
	mdh := data_handler.NewMemoryDataHandler(true, 10000000)
	ss := service.NewSenderService(c, data_handler.NewDataHandler(mdh))
	sc := client.NewSenderClient(ss)
	sc.StartSending()

}

func Register(cmd *cobra.Command) {
	var sender = &cobra.Command{
		Use:   "sender",
		Short: "sender client starts to send requests rapidly",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
	cmd.AddCommand(sender)
}
