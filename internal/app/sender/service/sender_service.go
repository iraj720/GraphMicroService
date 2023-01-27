package service

import (
	"context"
	"graph/pkg/data_handler"
	"graph/proto/reciever"
)

type SenderService interface {
	SendRequest() error
}

type senderService struct {
	rc reciever.RecieverClient
	dh data_handler.DataHandler
}

func NewSenderService(client reciever.RecieverClient, dataHandler data_handler.DataHandler) SenderService {

	return &senderService{rc: client, dh: dataHandler}
}

func (ss *senderService) SendRequest() error {
	gd, err := ss.dh.ReadData()
	if err != nil {
		return err
	}
	_, err = ss.rc.Send(context.Background(), &reciever.GraphDataRequest{Data: string(gd.Content)})
	if err != nil {
		// TODO : handling connection failures
		ss.dh.WriteData(gd)
		return err
	}
	return nil
}
