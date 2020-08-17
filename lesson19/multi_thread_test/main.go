package main

import (
	"fmt"
	"time"
)

func isPerm(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func test1(c chan int) {
	var i int
	for {
		i++
		result := isPerm(i)
		if !result {
			c <- i
		}
	}
}

func test2(c chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	var c chan int = make(chan int)
	go test1(c)
	go test2(c)

	time.Sleep(time.Hour)
}
