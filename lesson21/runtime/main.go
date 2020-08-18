package main

import (
	"fmt"
	"runtime"
	"time"
)

var i int

func cale() {
	for {
		i++
	}
}

func main() {
	cpu := runtime.NumCPU()
	fmt.Printf("cpu: %d\n", cpu)

	runtime.GOMAXPROCS(2)
	for i := 0; i < 10; i++ {
		go cale()
	}
	time.Sleep(time.Minute)
}
