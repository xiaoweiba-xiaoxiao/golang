package main

import "fmt"

func testArray1() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Printf("a=%d\n", a)
}

func testArray2() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a=%d\n", a)
}
func testArray3() {
	var a [3][2]int = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("a=%d\n", a)
}

func testArray4() {
	b := [5][2]int{4: {0: 20}}
	var a [3][2]int = [3][2]int{2: {1: 100}}
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)
}

func main() {
	// testArray1()
	// testArray2()
	//testArray3()
	testArray4()
}
