package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("the goroutine")
}

func main() {
	go hello()
	fmt.Println("the main")
	time.Sleep(time.Second)
}
