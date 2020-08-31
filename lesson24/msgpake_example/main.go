package main

import (
	"github.com/vmihailenco/msgpack"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func writejson() (err error) {
	var persons []*Person
	for i := 0; i < 10; i++ {
		person := &Person{
			Name: "hehe",
			Age:  18,
			Sex:  "女",
		}
		persons = append(persons, person)
	}
	data, ok := msgpack.Marshal(persons)
	if ok != nil {
		fmt.Println(err)
		return ok
	}
	err = ioutil.WriteFile("/home/golang/log/danmei.txt", data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func readjson() (err error) {
	var persons []Person
	data, ok := ioutil.ReadFile("/home/golang/log/danmei.txt")
	if ok != nil {
		fmt.Println(ok)
		return ok
	}
	err = msgpack.Unmarshal(data, &persons)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", persons)
	return
}

func main() {
	writejson()
	readjson()
}