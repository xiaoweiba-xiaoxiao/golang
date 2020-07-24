package main

import (
	"fmt"
	"os"
	"reflect"
)

type Float struct {
}

func reflectInterface(a interface{}) {
	t := reflect.TypeOf(a)
	k := t.Kind()
	fmt.Fprintln(os.Stdout, "type: ", t)
	fmt.Println(k)
}

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("type: ", v.Type()) //功能喝reflect.TypeOf一样
	k := v.Kind()
	switch k {
	case reflect.Float32:
		fmt.Printf("value is %f,type is %v\n", v.Float(), k)
	case reflect.Int:
		fmt.Printf("value is %d,type is %v\n", v.Int(), k)
	}
}

func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Ptr:
		v.Elem().SetInt(8)
	default:
	}
}

func main() {
	//var a int = 15
	var b int = 15
	reflect_value(b)
	reflect_set_value(&b)
	fmt.Println(b)
}
