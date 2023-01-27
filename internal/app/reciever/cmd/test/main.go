package main

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	// rc := controller.NewRecieverController()

	// lis, err := net.Listen(SERVER_TYPE, fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT))
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// reciever.RegisterRecieverServer(s, rc)
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }

	// fmt.Println("Server Running...")
	// server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	// if err != nil {
	// 	fmt.Println("Error listening:", err.Error())
	// 	os.Exit(1)
	// }
	// defer server.Close()
	// i := 0
	// for {
	// 	connection, err := server.Accept()
	// 	if err != nil {
	// 		fmt.Println("Error accepting: ", err.Error())
	// 		os.Exit(1)
	// Println()
	// 	fmt.Printf("client connected %d\n", i)
	// 	go processClient(connection)
	// 	i++
	// }
}

// func processClient(connection net.Conn) {
// 	buffer := make([]byte, 1024)
// 	mLen, err := connection.Read(buffer)
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}
// 	fmt.Println("Received: ", string(buffer[:mLen]))
// 	connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
// }
