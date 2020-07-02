package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name" db:"name" yaml:"name"`
	Age   int
	Sex   int
	Score float32
}

func (s *Student) SexToString() string {
	switch s.Sex {
	case 1:
		return "male"
	case 2:
		return "female"
	default:
		return "xiannv"
	}
}

func (s *Student) SetScore(score float32) {
	s.Score = score
}

func reflect_method(a interface{}) {
	v := reflect.ValueOf(a)
	t := v.Type()
	fmt.Printf("%s has method %d\n", t, t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		fmt.Printf("name=%s value=%s\n", t.Method(i).Name, t.Method(i).Type)
	}
	fmt.Printf("%#v\n", a)
	m := v.MethodByName("SexToString")
	var args []reflect.Value
	fmt.Println(m.Call(args))

	m1 := v.Method(0)
	var args1 []reflect.Value
	var score float32 = 98.5
	scorearg := reflect.ValueOf(score)
	args1 = append(args1, scorearg)
	m1.Call(args1)
	fmt.Printf("%#v\n", a)

}

func main() {
	var s Student = Student{
		Name: "xiaoxiao",
		Age:  29,
		Sex:  3,
	}
	reflect_method(&s)
	t := reflect.TypeOf(s)
	fmt.Println(t.Field(0).Tag.Get("json"))
}
