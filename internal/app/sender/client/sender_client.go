package client

import (
	"fmt"
	"graph/internal/app/sender/service"
	"math/rand"
	"time"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

type SenderClient struct {
	ss service.SenderService
}

func NewSenderClient(ss service.SenderService) SenderClient {
	return SenderClient{ss: ss}
}

func (sc *SenderClient) StartSending() {
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
					err := sc.ss.SendRequest(string(str))
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
