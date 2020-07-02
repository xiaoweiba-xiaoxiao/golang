package main

import (
	"fmt"
)

type testInterface interface {
}

func main() {
	var a int = 6
	var s string = "hello"
	var t testInterface
	t = a
	fmt.Println(t)
	t = s
	fmt.Println(t)
}
