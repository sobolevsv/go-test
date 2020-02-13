package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printAsync(arg string) {
	fmt.Println(arg)
	wg.Done()
}

func main() {

	wg.Add(2)

	go printAsync("hello")
	go printAsync("world")

	wg.Wait()
}
