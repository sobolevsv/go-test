package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello")
		wg.Done()
	}()

	go func() {
		fmt.Println("world")
		wg.Done()
	}()

	wg.Wait()
}
