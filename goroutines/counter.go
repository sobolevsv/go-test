package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int32 = 0

	for i := int32(0); i < 5; i++ {
		go func() {
			atomic.AddInt32(&counter, 1)
			for {
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Printf("counter: %v\n", counter)
}
