package main

import (
	"fmt"
	"time"
)

func inputChan(c chan int) {
	c <- 100
}

func outputChan(c chan int) {
	data := <-c
	fmt.Println(data)
}

func main() {
	var c chan int = make(chan int)
	go inputChan(c)
	go outputChan(c)
	time.Sleep(time.Second)
}
