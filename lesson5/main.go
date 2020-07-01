package main

import "fmt"

func testSlice1() {
	a := [5]int{1, 2, 3, 4, 5}
	// var b []int
	b := a[1:4]
	fmt.Printf("type of b is %v\n", b)
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])

}

func testSlice2() {
	b := []int{1, 2, 3, 4, 5}
	fmt.Printf("type of  b is %T\n", b)
}

func main() {
	//testSlice1()
	testSlice2()
}
