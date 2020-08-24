package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int32
var mutex sync.Mutex
var wg sync.WaitGroup

func addMutex() {

	for i := 0; i < 100000; i++ {
		mutex.Lock()
		x += 1
		mutex.Unlock()
	}
	wg.Done()
}

func add() {

	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&x, 1)
	}

	wg.Done()
}

func main() {
	start_time := time.Now().UnixNano()
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go addMutex()
	}
	wg.Wait()
	end_time := time.Now().UnixNano()
	fmt.Printf("addMutex is used %v\n", (end_time-start_time)/1000/1000)
	start_time = time.Now().UnixNano()
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	end_time = time.Now().UnixNano()
	fmt.Printf("add is used %v\n", (end_time-start_time)/1000/1000)
}
