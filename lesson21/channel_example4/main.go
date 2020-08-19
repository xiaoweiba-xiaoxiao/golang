package main

import (
	"fmt"
)

func dilter(chnl chan int, num int) {
	for num != 0 {
		dil := num % 10
		chnl <- dil
		num /= 10
	}
	close(chnl)
}
func squreFunc(squres chan int, num int) {
	chal := make(chan int)
	sum := 0
	go dilter(chal, num)
	for i := range chal {
		sum += i * i
	}
	squres <- sum
}

func cubeFunc(cube chan int, num int) {
	chal := make(chan int)
	sum := 0
	go dilter(chal, num)
	for i := range chal {
		sum += i * i * i
	}
	cube <- sum
}

func main() {
	num := 144
	squre := make(chan int)
	cube := make(chan int)
	go squreFunc(squre, num)
	go cubeFunc(cube, num)
	squres, cubes := <-squre, <-cube
	fmt.Println("Final output", squres+cubes)
}
