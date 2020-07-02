package lesson1

import (
	"fmt"
)

func init() {
	fmt.Println("lesson1 is init")
}

func testString() string {
	var str = "hello"
	var str2 = ""
	for i :=0; i < len(str);i++ {
		str2 += fmt.Sprintf("str[i] = %c\n",str[len(str)-i])	
	}
    return str2
}