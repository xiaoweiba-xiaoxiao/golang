package main

import (
	"fmt"
)

func defFunc(b int) {
	fmt.Println("input is ", b)
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic that is: ", r)
	}
}

func main() {
	var b int = 10
	defer defFunc(b)
	if b != 5 {
		panic("b is not 5")
	}
}
