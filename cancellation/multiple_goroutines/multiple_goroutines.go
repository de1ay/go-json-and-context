package multiplegoroutines

import (
	"context"
	"log"
	"time"
)

func goroutine(ctx context.Context, goroutineIndex int) {
	for {
		log.Printf("Goroutine #%d: sleeping for 1 minute...", goroutineIndex)
		select {
		case <-time.After(1 * time.Minute):
			continue
		case <-ctx.Done():
			log.Printf("Goroutine #%d: cancelled by context.", goroutineIndex)
			return
		}
	}
}

/*
	Example of execution logs:

	2021/12/28 22:57:14 Goroutine #3: sleeping for 1 minute...
	2021/12/28 22:57:14 Goroutine #1: sleeping for 1 minute...
	2021/12/28 22:57:14 Goroutine #2: sleeping for 1 minute...
	2021/12/28 22:57:19 Goroutine #2: cancelled by context.
	2021/12/28 22:57:19 Goroutine #3: cancelled by context.
	2021/12/28 22:57:19 Goroutine #1: cancelled by context.
*/

func Exec() {
	bCtx := context.Background()
	timeoutConext, cancel := context.WithTimeout(bCtx, 5*time.Second)
	defer cancel()

	go goroutine(timeoutConext, 1)
	go goroutine(timeoutConext, 2)
	go goroutine(timeoutConext, 3)

	<-timeoutConext.Done()

	// wait for gouroutines to stop
	time.Sleep(1 * time.Second)
}
