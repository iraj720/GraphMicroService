package reciever

import (
	"fmt"
	"graph/internal/app/reciever/controller"
	"graph/internal/app/reciever/service"
	"graph/pkg/data_handler"
	socketclient "graph/pkg/socket_client"
	"graph/proto/reciever"
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
	CLIENT_HOST = "localhost"
	CLIENT_PORT = "9989"
)

var count = 0

func main() {
	sc := socketclient.NewSocketClient(CLIENT_HOST, CLIENT_PORT, SERVER_TYPE)
	mdh := data_handler.NewMemoryDataHandler(false, 0)
	dh := data_handler.NewDataHandler(mdh)
	rs := service.NewRecieverService(sc, dh)
	rc := controller.NewRecieverController(rs)
	rc.StartSending()
	rs.StartHandlingFailedRequests()
	lis, err := net.Listen(SERVER_TYPE, fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reciever.RegisterRecieverServer(s, rc)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func Register(cmd *cobra.Command) {
	var reciever = &cobra.Command{
		Use:   "reciever",
		Short: "reciever recieves from sender and sends to broker. its client will start sending requests whenever they achieved",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
	cmd.AddCommand(reciever)
}
