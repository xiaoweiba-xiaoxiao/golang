package main

import (
	"fmt"
)

func febonacii(c chan int) {
	a, b, d := 0, 1, 0
	for i := 0; i < 10; i++ {
		d = a
		a, b = b, a+b
		c <- d
	}
	close(c)
}

func main() {
	var c chan int = make(chan int)
	go febonacii(c)
	for i := range c {
		fmt.Println(i)
	}
}
