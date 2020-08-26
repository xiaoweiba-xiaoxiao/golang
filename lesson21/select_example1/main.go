package main

import (
	"fmt"
	"time"
)

func server1(chnl chan string) {
	time.Sleep(time.Second * 6)
	chnl <- "from server1"
}

func server2(chnl chan string) {
	time.Sleep(time.Second * 3)
	chnl <- "from server2"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go server1(ch1)
	go server2(ch2)
	select {
	case data := <-ch1:
		fmt.Println(data)
	case data := <-ch2:
		fmt.Println(data)
	}
}
