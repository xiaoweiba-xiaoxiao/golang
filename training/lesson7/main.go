package main

import "fmt"

func test1() {
	var b int = 100
	var a *int = &b
	fmt.Printf("b adress is %p\n", a)
}

func modifyTest(a *int) {
	*a = 100
}

func test2() {
	b := 300
	a := &b
	modifyTest(a)
	fmt.Printf("b=%d", b)
}

func main() {
	test1()
	test2()
}
