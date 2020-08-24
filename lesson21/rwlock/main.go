package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var rwmutex sync.RWMutex
var wg sync.WaitGroup

func read(i int) {
	rwmutex.RLock()
	fmt.Println(i, "is read")
	time.Sleep(time.Second)
	rwmutex.RUnlock()
	wg.Done()
}

func write() {
	rwmutex.Lock()
	x += 1
	fmt.Printf("%d is write\n", x)
	time.Sleep(time.Second)
	rwmutex.Unlock()
	wg.Done()
}

func main() {
	wg.Add(1)
	go write()
	time.Sleep(time.Second * 2)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read(i)
	}
	wg.Wait()
}
