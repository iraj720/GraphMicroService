package service

import (
	"fmt"
	"graph/pkg/data_handler"
	socketclient "graph/pkg/socket_client"
	"time"
)

type BrokerService interface {
	SendRequest(*data_handler.GraphData) error
	StartHandlingFailedRequests()
}

type brokerService struct {
	sc socketclient.SocketClient
	dh data_handler.DataHandler
}

func NewBrokerService(sc socketclient.SocketClient, dh data_handler.DataHandler) BrokerService {
	return &brokerService{sc: sc, dh: dh}
}

// this will run forever
func (bs *brokerService) StartHandlingFailedRequests() {
	go func() {
		sleepTime := 2
		for {
			if bs.dh.Size() > 0 {
				gd, err := bs.dh.ReadData()
				if err != nil {
					err := bs.dh.WriteData(gd)
					if err != nil {
						fmt.Printf("unable to handle failed data for '%d' seconds\n", sleepTime)
					}
					time.Sleep(time.Duration(sleepTime) * time.Second)
					continue
				}
				err = bs.sc.SendRequest(gd)
				if err != nil {
					bs.dh.WriteData(gd)
				}
			}
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
	}()

}

func (bs *brokerService) SendRequest(gd *data_handler.GraphData) error {
	err := bs.sc.SendRequest(gd)
	if err != nil {
		gd.TransferState = data_handler.TransferState_FAILED
		bs.dh.WriteData(gd)
		return err
	}
	return nil
}
