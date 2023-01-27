package broker

import (
	"graph/internal/app/broker_app/broker/controller"
	"graph/internal/app/broker_app/broker/service"
	"graph/internal/app/broker_app/logger"
	"graph/pkg/data_handler"
	socketclient "graph/pkg/socket_client"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9989"
	SERVER_TYPE = "tcp"
	CLIENT_HOST = "localhost"
	CLIENT_PORT = "9990"
	CLIENT_TYPE = "tcp"
)

func main() {
	sc := socketclient.NewSocketClient(CLIENT_HOST, CLIENT_PORT, CLIENT_TYPE)
	mdh := data_handler.NewMemoryDataHandler(false, 0)
	dh := data_handler.NewDataHandler(mdh)
	bs := service.NewBrokerService(sc, dh)
	f, err := os.OpenFile("broker_app_logger.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	loger := log.New()
	loger.Out = f
	loger.Formatter = &log.JSONFormatter{}
	defer f.Close()
	bc := controller.NewBrokerController(logger.NewGraphLogger(loger), bs)
	bs.StartHandlingFailedRequests()

	err = bc.StartServing(SERVER_HOST, SERVER_PORT, SERVER_TYPE)
	if err != nil {
		logrus.Fatalf("cannot start serving %v", err)
	}
}

func Register(root *cobra.Command) {
	var broker = &cobra.Command{
		Use:   "broker",
		Short: "broker recieves messages, log them and send them to destination",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
	root.AddCommand(broker)
}
