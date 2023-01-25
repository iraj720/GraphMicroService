package socketclient

import (
	"fmt"
	"net"
	"time"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func Connect() {
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	for {
		_, err = connection.Write([]byte("Hello Server! Greetings."))
		//buffer := make([]byte, 1024)
		//mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		//fmt.Println("Received: ", string(buffer[:mLen]))
		time.Sleep(5 * time.Second)
	}

}
