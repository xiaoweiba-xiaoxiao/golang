package main

import (
	"fmt"
)

func putChan(chnl chan int) {
	chnl <- 100
}

func main() {
	var c chan int = make(chan int)
	go putChan(c)
	data := <-c
	fmt.Println(data)
}
