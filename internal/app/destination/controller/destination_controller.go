package controller

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type destinationController struct {
}

func NewDestinationController() *destinationController {
	return &destinationController{}
}

func (bc *destinationController) StartServing(ServerHost string, ServerPort string, ServerType string) {
	fmt.Println("Server Running... On ", ServerHost, ":", ServerPort)
	server, err := net.Listen(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	t1 := time.Now()
	defer server.Close()
	count := 0
	defer func() {
		var s chan os.Signal
		signal.Notify(s, syscall.SIGTERM)
		<-s
		fmt.Printf("number of recieved messages : %d", count)
	}()
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
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}

		//fmt.Printf("client connected %d\n", i)
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
