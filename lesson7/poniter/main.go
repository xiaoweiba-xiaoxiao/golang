package main

import "fmt"

func main() {
	// testPoint1()
	//testPoint3()
	// testPoint4()
	// testPoint5()
	// testPoint6()
	testPoint7()
}

func testPoint1() {
	var b *int = new(int)
	*b = 100
	fmt.Printf("b=%d\n", *b)
}

func testPoint2() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
}

func testPoint3() {
	a := 25
	var b *int
	if b == nil {
		fmt.Printf("b is %p\n", b)
		b = &a
		fmt.Println("b after initialization is", b)
	}
}

func modifySlice(a *[]int) {
	(*a)[0] = 100
}

func testPoint4() {
	var a *[]int = &[]int{4, 5, 6, 7}
	fmt.Println(*a)
	modifySlice(a)
	fmt.Println(*a)
}

func testPoint5() {
	var a *int = new(int)
	*a = 100
	fmt.Printf("*a=%d\n", *a)
	var b *[]int = new([]int)
	//(*b)[0]=100 错误操作，b是一个空切片
	*b = make([]int, 5, 100) //使用make初始化*b
	(*b)[0] = 100
	(*b)[1] = 200
	fmt.Printf("*b=%d\n", *b)
}

func modifyInt(a int) {
	a = 100
}

func testPoint6() {
	b := 10
	modifyInt(b)
	fmt.Println(b)
}

func testPoint7() {
	var a int = 100
	var b *int = &a
	var c *int = b
	*c = 200
	fmt.Printf("*c=%d,*b=%d,a=%d\n", *c, *b, a)
}
