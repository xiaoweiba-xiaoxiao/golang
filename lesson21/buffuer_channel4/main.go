package main

import (
	"fmt"
	"sync"
	"time"
)

func process(wg *sync.WaitGroup) {
	fmt.Println("prcess begin")
	time.Sleep(time.Second)
	fmt.Println("process end")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go process(&wg)
	}
	wg.Wait()
}
