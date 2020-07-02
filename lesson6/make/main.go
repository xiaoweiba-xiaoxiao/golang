package main

import "fmt"

func testMake1() {
	a := make([]int, 0, 0)
	fmt.Printf("a.len=%d,a.caple=%d\n", len(a), cap(a))
	a = append(a, 11)
	fmt.Printf("a=%v\n", a)
	fmt.Printf("a.len=%d,a.caple=%d\n", len(a), cap(a))
}
func testCap() {
	a := []string{"a", "b", "c", "d", "e", "f", "g"}
	b := a[1:4]
	fmt.Printf("b=%s,b.len=%d,b.cap=%d\n", b, len(b), cap(b))
	b = b[:cap(b)]
	fmt.Printf("b=%s,b.len=%d,b.cap=%d\n", b, len(b), cap(b))
}

func testSlice() {
	var a []int
	fmt.Printf("%p,len(a)=%d,cap(a)=%d\n", a, len(a), cap(a))
	if a == nil {
		fmt.Println("a is nil")
	}
	a = append(a, 100)
	fmt.Printf("%p,len(a)=%d,cap(a)=%d\n", a, len(a), cap(a))
	a = append(a, 200)
	fmt.Printf("%p,len(a)=%d,cap(a)=%d\n", a, len(a), cap(a))
	a = append(a, 300)
	fmt.Printf("%p,len(a)=%d,cap(a)=%d\n", a, len(a), cap(a))
	a = append(a, 400)
	fmt.Printf("%p,len(a)=%d,cap(a)=%d\n", a, len(a), cap(a))
	a = append(a, 500)
	fmt.Printf("%p,len(a)=%d,cap(a)=%d\n", a, len(a), cap(a))
	a = append(a, 600)
	fmt.Printf("%p,len(a)=%d,cap(a)=%d\n", a, len(a), cap(a))
}

func testAppend() {
	var a []int = []int{1, 2, 3}
	var b []int = []int{4, 5, 6}
	fmt.Printf("a=%d,a.len=%d,a.cap=%d\n", a, len(a), len(a))
	a = append(a, b...)
	fmt.Printf("a=%d,a.len=%d,a.cap=%d\n", a, len(a), len(a))
}

func main() {
	// testMake1()
	//testCap()
	// testSlice()
	testAppend()
}
