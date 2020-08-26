package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex

func add(x *int) {
	mutex.Lock()
	*x = *x + 1
	mutex.Unlock()
	(&wg).Done()
}

func main() {
	x := 0
	wg.Add(2)
	go add(&x)
	go add(&x)
	wg.Wait()
	fmt.Println(x)

}
