package sender

import (
	"fmt"
	"graph/internal/app/sender/service"
	"math/rand"
	"time"
)

type SenderClient struct {
	ss service.SenderService
}

func NewSenderClient(ss service.SenderService) SenderClient {
	return SenderClient{ss: ss}
}

func (sc *SenderClient) StartSending() {
	max := 50
	min := 8
	count := 0
	gosStarted := 0

	t1 := time.Now()
	for i := 0; i < 100; i++ {
		go func() {
			for time.Since(t1).Seconds() <= 2 {
				number := rand.Int31n(int32(max-min)) + int32(min)
				err := sc.ss.SendRequest(fmt.Sprintf("%d", number))
				if err != nil {
					fmt.Println(err)
				} else {
					count++
				}
			}

		}()
	}
	fmt.Println(time.Since(t1), " ", gosStarted)
	for time.Since(t1).Seconds() <= 2 {

	}
	fmt.Println(time.Since(t1), " gos ", gosStarted, " count ", count)
}
