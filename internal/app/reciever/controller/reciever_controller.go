package controller

import (
	"context"
	"fmt"
	"graph/proto/reciever"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type RecieverController struct {
	reciever.UnimplementedRecieverServer
}

func NewRecieverController() *RecieverController {
	return &RecieverController{}
}

func (rc *RecieverController) Send(ctx context.Context, in *reciever.GraphDataRequest) (*reciever.GraphDataResponse, error) {
	log.Printf("Received: %v", in.Data)
	return &reciever.GraphDataResponse{Message: "Hello "}, nil
}

func (rc *RecieverController) SendData(c echo.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in %f \n", r)
		}
	}()
	headerVal := c.Request().Header.Get("Content-Type")
	if headerVal != "application/octet-stream" {
		return c.JSON(http.StatusBadRequest, fmt.Errorf("content-type is not acceptable"))
	}
	var data []byte
	_, err := c.Request().Body.Read(data)
	fmt.Println("this is data : ", data)
	// c.Response().Header().Set("Content-Disposition", "attachment; filename="+fileName)
	// c.Response().Header().Set("Content-Type", "application/octet-stream")
	// _, err = c.Response().Write(b)
	return err
}
