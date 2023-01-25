package main

import (
	"fmt"
	"graph/internal/app/reciever/controller"
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
)

var count = 0

func server() {
	rc := controller.NewRecieverController()

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

func client() {

}

func Register(cmd *cobra.Command) {
	var reciever = &cobra.Command{
		Use:       "reciever",
		Short:     "reciever recieves from sender and sends to broker. its client will start sending requests whenever they achieved",
		ValidArgs: []string{"server", "client"},
		Args:      cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			serverOrClient := args[0]
			if serverOrClient == "server" {
				server()
			} else if serverOrClient == "client" {
				client()
			}
		},
	}
	cmd.AddCommand(reciever)
}
