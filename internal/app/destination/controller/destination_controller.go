package controller

import (
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

type destinationController struct {
}

func NewDestinationController() *destinationController {
	return &destinationController{}
}

func (dc *destinationController) StartServing(ServerHost string, ServerPort string, ServerType string) error {
	fmt.Println("Server Running... On ", ServerHost, ":", ServerPort)
	server, err := net.Listen(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		return err
	}
	defer server.Close()

	t1 := time.Now()
	count := 0
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println(time.Since(t1), " count ", count, " req/s ", count/int(time.Since(t1).Seconds()))
		}
	}()
	buffer := make([]byte, 100000)
	for {
		connection, err := server.Accept()
		if err != nil {
			logrus.Info("cannot accept connection")
			continue
		}
		for {
			_, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
				break
			}
			count++
		}
	}
}
