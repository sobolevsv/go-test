package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(arg string) {
		fmt.Println(arg)
		wg.Done()
	}("hello")

	go func(arg string) {
		fmt.Println(arg)
		wg.Done()
	}("world")

	wg.Wait()
}
