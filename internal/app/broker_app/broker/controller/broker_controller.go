package controller

import (
	"fmt"
	"graph/internal/app/broker_app/broker/service"
	"graph/internal/app/broker_app/logger"
	"graph/pkg/data_handler"
	"net"
)

type brokerController struct {
	logger        logger.GraphLogger
	brokerService service.BrokerService
}

func NewBrokerController(logger logger.GraphLogger, brokerService service.BrokerService) *brokerController {
	return &brokerController{logger: logger, brokerService: brokerService}
}

func (bc *brokerController) StartServing(ServerHost string, ServerPort string, ServerType string) error {
	fmt.Println("Server Running...")
	server, err := net.Listen(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		return err
	}
	defer server.Close()

	buffer := make([]byte, 100000)
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		for {
			mLen, err := connection.Read(buffer)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
				break
			}
			bc.logger.Log(string(buffer[:mLen]))
			bc.brokerService.SendRequest(&data_handler.GraphData{Content: buffer[:mLen]})
			if err != nil {
				fmt.Println("request failed")
			}
		}
	}
}
