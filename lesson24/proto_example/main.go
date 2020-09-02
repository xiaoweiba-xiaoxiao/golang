package main

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"xiao.com/golang/lesson/lesson24/proto_example/address"
)

func marshal() {
	var phonebook address.PhoneBook
	for i := 0; i < 64; i++ {
		person := &address.Person{
			Id:   int32(i),
			Name: fmt.Sprintf("danmei%d", i),
		}

		phone := &address.Phone{
			Type:   address.PhoneType_HOME,
			Number: "18888888888",
		}
		person.Phones = append(person.Phones, phone)
		phone = &address.Phone{
			Type:   address.PhoneType_Work,
			Number: "028888888",
		}
		person.Phones = append(person.Phones, phone)
		phonebook.Persons = append(phonebook.Persons, person)
	}
	data, err := proto.Marshal(&phonebook)
	if err != nil {
		fmt.Println("proto marshal failed,error:", err)
		return
	}
	err = ioutil.WriteFile("/home/golang/log/test.pro", data, 0755)
	if err != nil {
		fmt.Println("write faild,error", err)
		return
	}
}

func unmarshal() {
	var phonebook address.PhoneBook
	data, err := ioutil.ReadFile("/home/golang/log/test.pro")
	if err != nil {
		fmt.Println("read faild,error:", err)
		return
	}
	err = proto.Unmarshal(data, &phonebook)
	if err != nil {
		fmt.Println("unmarshal faild,error:", err)
		return
	}
	fmt.Printf("%v", phonebook)
}

func main() {
	unmarshal()
}
