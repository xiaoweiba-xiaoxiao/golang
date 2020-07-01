package main

import "fmt"

func modify(a *int) {
	*a = 100
}

func modifyarr(a *[]int) {
	(*a)[0] = 100
}

func testpointer() {
	a := []int{1, 2, 3, 5, 6}
	fmt.Printf("b:%d\n", a)
	modifyarr(&a)
	fmt.Printf("b:%d\n", a)
}

func testexample() {
	var b int = 10
	fmt.Printf("b:%d\n", b)
	p := &b
	modify(p)
	fmt.Printf("b:%d\n", b)
}

func main() {
	// testexample()
	testpointer()
}
