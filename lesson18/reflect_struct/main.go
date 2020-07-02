package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Sex  int
	Age  int
}

func reflect_struct(a interface{}) {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	k := v.Kind()
	fmt.Println(t, v, k)
}

func reflect_struct_filed(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Float32:
		fmt.Println("a is float32")
	case reflect.Int64:
		fmt.Println("a is int64")
	case reflect.Struct:
		fmt.Println("a is struct")
		fmt.Println(t.NumField())
		for i := 0; i < v.NumField(); i++ {
			filed := v.Field(i)
			fmt.Printf("name:%s,type:%s,value:%v\n", t.Field(i).Name, filed.Type().Kind(), filed.Interface())
		}
	default:
		fmt.Println("unkown type of a")
	}
}

func reflect_struct_set(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	k := t.Kind()
	switch k {
	case reflect.Ptr:
		v.Elem().FieldByName("Name").SetString("xiaoxiao")
		v.Elem().Field(1).SetInt(2)
		v.Elem().FieldByName("Age").SetInt(29)
	}
}

func main() {
	var a *Student = &Student{
		// Name: "xiaoxiao",
		// Sex:  2,
		// Age:  29,
	}
	// reflect_struct(a)
	reflect_struct_set(a)
	fmt.Println(a)
}
