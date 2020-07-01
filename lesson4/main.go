package main

import (
	"fmt"
	"strings"
	"time"
)

func Adder() func(int) int {
	var x int
	return func(d int) int {
		x += d
		return x
	}
}

func Adder1(base int) func(int) int {
	return func(d int) int {
		base += d
		return base
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func testClosure3() {
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test"))
	fmt.Println(func2("test"))
}

func testClosure2() {
	f := Adder1(10)
	fmt.Println(f(1), f(2))
	f1 := Adder1(100)
	fmt.Println(f1(1), f1(2))
}

func testClosure1() {
	f := Adder()
	ret := f(1)
	fmt.Printf("ret=%d\n", ret)
	ret = f(20)
	fmt.Printf("ret=%d\n", ret)
	ret = f(100)
	fmt.Printf("ret=%d\n", ret)
	f1 := Adder()
	ret = f1(1)
	fmt.Printf("ret=%d\n", ret)
}

func casle(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(b int) int {
		base -= b
		return base
	}
	return add, sub
}

func testClosure4() {
	f1, f2 := casle(10)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
	fmt.Println(f1(7), f2(8))
}

func testClosure5() {
	for i := 0; i < 5; i++ {
		/* 闭包的副作用，绑定i,全部打印5
		go func() {
			fmt.Println(i)
		}()
		*/
		go func(index int) {
			fmt.Println(index)
		}(i)
		time.Sleep(time.Second)
	}
}

func main() {
	// testClosure1()
	//testClosure2()
	//testClosure3()
	//testClosure4()
	testClosure5()
}
