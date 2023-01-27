package client

import (
	"fmt"
	"graph/internal/app/sender/service"
	"sync"
	"time"
)

type SenderClient struct {
	ss service.SenderService
}

func NewSenderClient(ss service.SenderService) SenderClient {
	return SenderClient{ss: ss}
}

func (sc *SenderClient) StartSending() {

	count := 0
	gosStarted := 0
	gosNeeded := 100

	t1 := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < gosNeeded; i++ {
		gosStarted++
		wg.Add(1)

		defer wg.Done()
		go func() {
			// runtime.LockOSThread()
			// defer runtime.UnlockOSThread()
			// data := "hello"

			for {
				err := sc.ss.SendRequest()
				if err != nil {
					fmt.Println(err)
				} else {
					count++
				}
				//time.Sleep(100 * time.Millisecond)
			}

		}()

	}
	go func() {
		for {
			time.Sleep(2 * time.Second)
			fmt.Println(time.Since(t1), " goroutines ", gosStarted, " count ", count, " req/s ", count/int(time.Since(t1).Seconds()))
		}
	}()
	wg.Wait()
}
