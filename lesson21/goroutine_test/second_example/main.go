package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello this is %d", i)
}

func 

func main() {
	for i := 0; i < 10; i++ {
		go hello(i)
	}
	time.Sleep(time.Second)
}