package service

import (
	"fmt"
	"graph/pkg/data_handler"
	socketclient "graph/pkg/socket_client"
	"time"
)

type RecieverService interface {
	SendRequest(*data_handler.GraphData) error
	StartHandlingFailedRequests()
}

type recieverService struct {
	sc socketclient.SocketClient
	dh data_handler.DataHandler
}

func NewRecieverService(sc socketclient.SocketClient, dh data_handler.DataHandler) RecieverService {
	return &recieverService{sc: sc, dh: dh}
}

func (ss *recieverService) SendRequest(gd *data_handler.GraphData) error {
	err := ss.sc.SendRequest(gd)
	if err != nil {
		gd.TransferState = data_handler.TransferState_FAILED
		ss.dh.WriteData(gd)
		return err
	}
	return nil
}

// this will run forever
func (ss *recieverService) StartHandlingFailedRequests() {
	go func() {
		sleepTime := 2
		for {
			if ss.dh.Size() > 0 {
				gd, err := ss.dh.ReadData()
				if err != nil {
					err := ss.dh.WriteData(gd)
					if err != nil {
						fmt.Printf("unable to handle failed data for '%d' seconds\n", sleepTime)
					}
					time.Sleep(time.Duration(sleepTime) * time.Second)
					continue
				}
				err = ss.sc.SendRequest(gd)
				if err != nil {
					ss.dh.WriteData(gd)
				}
			}
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
	}()

}
