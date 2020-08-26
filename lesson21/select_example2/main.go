package main

import (
	"fmt"
	"time"
)

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("wirte sucessful")
		default:
			fmt.Println("ch  is full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	ch := make(chan string, 10)
	go write(ch)
	for i := range ch {
		fmt.Printf("recv: value %s\n", i)
		time.Sleep(time.Second)
	}
}
