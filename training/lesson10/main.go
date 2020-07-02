package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
)

func countWorld(s string) (count map[string]int) {
	arr := strings.FieldsFunc(s, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	count = make(map[string]int, 1024)
	for _, key := range arr {
		count[key] += 1
	}
	return
}

func createStudent() (student map[int]map[string]interface{}) {
	student = make(map[int]map[string]interface{}, 128)
	// value,ok := student[id]
	// if !ok {
	// 	value = make(map[string]interface{},128)
	// }
	// value["age"] = age
	// value["sorce"] = sorce
	for i := 1; i < 100; i++ {
		value, ok := student[i]
		if !ok {
			value = make(map[string]interface{}, 128)
		}
		value["age"] = rand.Intn(23)
		value["sorce"] = rand.Intn(100)
		student[i] = value
	}
	return
}

func test2() {
	rand.Seed(time.Now().UnixNano())
	student := createStudent()
	fmt.Printf("%#v", student)
}

func test1() {
	s := "how do you do,I am fine,thank you,and you."
	fmt.Println(countWorld(s))
}

func main() {
	// test1()
	test2()
}
