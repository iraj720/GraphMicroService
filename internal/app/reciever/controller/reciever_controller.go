package controller

import (
	"context"
	"graph/internal/app/reciever/service"
	"graph/pkg/data_handler"
	"graph/proto/reciever"
	"log"
)

type RecieverController struct {
	reciever.UnimplementedRecieverServer
	ss service.RecieverService
	ch chan string
}

func NewRecieverController(socketService service.RecieverService) *RecieverController {
	return &RecieverController{ss: socketService}
}

func (rc *RecieverController) Send(ctx context.Context, in *reciever.GraphDataRequest) (*reciever.GraphDataResponse, error) {
	log.Printf("Received: %v", in.Data)
	go rc.ss.SendRequest(&data_handler.GraphData{Content: []byte(in.Data)})
	return &reciever.GraphDataResponse{Message: "Hello"}, nil
}
