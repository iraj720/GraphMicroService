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
	return &RecieverController{ss: socketService, ch: make(chan string, 1000)}
}

func (rc *RecieverController) Send(ctx context.Context, in *reciever.GraphDataRequest) (*reciever.GraphDataResponse, error) {
	log.Printf("Received: %v", in.Data)
	go rc.ss.SendRequest(&data_handler.GraphData{Content: []byte(in.Data)})
	//rc.ch <- in.Data
	return &reciever.GraphDataResponse{Message: "Hello"}, nil
}

func (rc *RecieverController) StartSending() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				data := <-rc.ch
				rc.ss.SendRequest(&data_handler.GraphData{Content: []byte(data)})
			}
		}()
	}
}
