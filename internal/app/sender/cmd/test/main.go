package main

import (
	"context"
	"fmt"
	"graph/proto/reciever"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", SERVER_HOST, SERVER_PORT), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := reciever.NewRecieverClient(conn)

	max := 50000
	min := 8
	count := 0
	gosStarted := 0

	t1 := time.Now()
	endTimeChannel := make(chan int)
	for i := 0; i < 1; i++ {
		go func() {
			timeCh := time.After(time.Second * 2)
		l2:
			for {
				select {
				case <-timeCh:
					endTimeChannel <- 1
					break l2
				default:
					number := rand.Int31n(int32(max-min)) + int32(min)
					// 8 B - 5 KB data
					str := make([]byte, number)
					_, err := c.Send(context.Background(), &reciever.GraphDataRequest{Data: string(str)})
					if err != nil {
						fmt.Println(err)
					} else {
						count++
					}
				}

			}

		}()
	}
	fmt.Println(time.Since(t1), " ", gosStarted)
	<-endTimeChannel
	fmt.Println(time.Since(t1), " gos ", gosStarted, " count ", count)
}
