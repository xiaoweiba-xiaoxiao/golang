package main

import (
	"fmt"
)

/* 设置返回值 -> defer函数执行 -> return */

//设置返回值为x,执行defer x + 1，return 6
func testDefer1() (x int) {
	x = 5
	defer func() {
		x += 1
	}()
	return
}

//设置x 为返回值，此时值为5 ，执行defer x = 6 ,return 5
func testDefer2() int {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func testDefer3() (y int) {
	x := 5
	defer func() {
		x += 1
	}()
	return x
}

func main() {
	// fmt.Println(testDefer1())
	fmt.Println(testDefer2())
}
