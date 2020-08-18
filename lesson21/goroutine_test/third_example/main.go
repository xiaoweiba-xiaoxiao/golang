package main

import (
	"fmt"
	"time"
)

func printInt() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printChar() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go printInt()
	go printChar()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println()
	fmt.Println("main terminated")
}
