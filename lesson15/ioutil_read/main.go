package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	fmt.Println(string(content))
}
