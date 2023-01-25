package service

import (
	"context"
	"fmt"
	"graph/proto/reciever"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

type SenderService interface {
	Healthz() (interface{}, error)
	SendRequest(data string) error
}

type senderService struct {
	client reciever.RecieverClient
}

func NewSenderService(client reciever.RecieverClient) SenderService {

	return &senderService{client: client}
}

func (ss *senderService) Healthz() (interface{}, error) {
	return "", fmt.Errorf("")
}

func (ss *senderService) SendRequest(data string) error {
	_, err := ss.client.Send(context.Background(), &reciever.GraphDataRequest{Data: data})
	if err != nil {
		return err
	}
	return nil
}
