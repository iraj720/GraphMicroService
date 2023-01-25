package main

import (
	"fmt"
	"graph/internal/app/reciever/controller"
	"graph/proto/reciever"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

var count = 0

func main() {
	rc := controller.NewRecieverController()

	// e := echo.New()
	// e.GET("/healthz", func(c echo.Context) error {
	// 	count++
	// 	fmt.Println(count)
	// 	return c.String(http.StatusOK, "Im UP")
	// })
	// e.Use(middleware.Recover())

	// rc.RegisterRoutes(e)

	// err := e.Start(fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
	// if err != nil {
	// 	logrus.Fatalf("unable to unable to run web server: %v", err)
	// }
	// lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// reciever.RegisterRecieverServer(rc, reciever.UnimplementedRecieverServer{})

	// grpcServer := grpc.NewServer()

	// reiever.RegisterChatServiceServer(grpcServer, &s)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %s", err)
	// }
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
