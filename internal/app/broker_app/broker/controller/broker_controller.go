package controller

import (
	"fmt"
	"graph/internal/app/broker_app/broker/service"
	"graph/internal/app/broker_app/logger"
	"graph/pkg/data_handler"
	"net"
	"os"
)

type brokerController struct {
	logger        logger.GraphLogger
	brokerService service.BrokerService
	buffer        []byte
}

func NewBrokerController(logger logger.GraphLogger, brokerService service.BrokerService) *brokerController {
	return &brokerController{logger: logger, brokerService: brokerService, buffer: make([]byte, 100000)}
}

func (bc *brokerController) StartServing(ServerHost string, ServerPort string, ServerType string) {
	fmt.Println("Server Running...")
	server, err := net.Listen(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		for {
			mLen, err := connection.Read(bc.buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
				break
			}
			bc.logger.Log(string(bc.buffer[:mLen]))
			bc.brokerService.SendRequest(&data_handler.GraphData{Content: bc.buffer[:mLen]})
			if err != nil {
				fmt.Println("request failed")
			}
		}
	}
}