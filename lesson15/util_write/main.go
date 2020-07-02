package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	str := "hello xiaoxiao!\n"
	for i := 0; i < 10; i++ {
		err := ioutil.WriteFile("test.txt", []byte(str), 0666)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
	}
}
