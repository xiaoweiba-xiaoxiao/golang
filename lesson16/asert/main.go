package main

import (
	"fmt"
)

type testInerface interface {
}

func Asert(a testInerface) {
	switch v := a.(type) {
	case int:
		fmt.Printf("%T,value=%d\n", v, v)
	case string:
		fmt.Printf("%T,value=%s\n", v, v)
	case float32:
		fmt.Printf("%T,value=%f\n", v, v)
	default:
		fmt.Println("can't asert the type")
	}
}

func main() {
	var a int = 34
	Asert(a)
	var s string = "hello"
	Asert(s)
	var f float32 = 3.12
	Asert(f)
	var d float64 = 3.1415
	Asert(d)
}
