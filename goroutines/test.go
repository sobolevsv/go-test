package main

import (
	"fmt"
	"sync"
)

func testGo(wg *sync.WaitGroup) {
	//    fmt.Print("Hello world\n")
	wg.Done()
	fmt.Print("exiting GOROUTINE !!!!!!\n")
}

func main() {
	var wg sync.WaitGroup

	//    for {
	wg.Add(1)

	fmt.Print("starting testGo\n")
	go testGo(&wg)

	wg.Wait()
	//    }

	fmt.Print("exiting MAIN!!!!!!!!!!!!!\n")
}
