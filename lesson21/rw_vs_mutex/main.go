package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var mutex sync.Mutex
var rwmutex sync.RWMutex
var wg sync.WaitGroup

func testMutex() {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		fmt.Printf("%d", x)
		mutex.Unlock()
	}
	wg.Done()
}

func testRWmutex() {
	for i := 0; i < 100000; i++ {
		rwmutex.RLock()
		fmt.Printf("%d", x)
		rwmutex.RUnlock()
	}
	wg.Done()
}

func main() {
	start_time := time.Now().UnixNano()
	wg.Add(1)
	go testMutex()
	end_time := time.Now().UnixNano()
	pay := end_time - start_time
	wg.Wait()
	fmt.Println()
	fmt.Println(pay)
	start_time = time.Now().UnixNano()
	wg.Add(1)
	go testRWmutex()
	end_time = time.Now().UnixNano()
	pay = end_time - start_time
	wg.Wait()
	fmt.Println()
	fmt.Println(pay)

}
