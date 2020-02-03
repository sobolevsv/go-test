package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var counter int64 = 0
	for i := 0; i < 10; i++ {
		go func(id int) {
			for {
				atomic.AddInt64(&counter, 1)
			}
		}(i)
	}
	for {
		if atomic.LoadInt64(&counter) >= 10000 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("counter: %v\n", counter)
}
