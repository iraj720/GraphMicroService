package socketclient

import (
	"graph/pkg/data_handler"
	"net"
)

type SocketClient interface {
	SendRequest(*data_handler.GraphData) error
}

type socketClient struct {
	ServerHost string
	ServerPort string
	ServerType string
	connection net.Conn
}

func NewSocketClient(ServerHost string, ServerPort string, ServerType string) SocketClient {
	connection, _ := net.Dial(ServerType, ServerHost+":"+ServerPort)
	// if err != nil {
	// 	return err
	// }
	//connection.Close()
	return &socketClient{ServerHost: ServerHost, ServerPort: ServerPort, ServerType: ServerType, connection: connection}
}

func (sc *socketClient) SendRequest(gd *data_handler.GraphData) error {

	n, err := sc.connection.Write(gd.Content)
	if n < len(gd.Content) || err != nil {
		return err
	}
	return nil
}
