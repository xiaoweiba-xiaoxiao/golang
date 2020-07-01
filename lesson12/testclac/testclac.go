package main

import (
	"fmt"
	"localhost/golang/geometry/lesson12/clac"
)

func testClac() {
	var user *clac.TestUser = clac.New("xiaoxiao", "girl", "wuhan", 29) //通过构造函数初始化一个Testuser结构体
	userinstrction := user.Say()
	fmt.Printf("%s\n", userinstrction)
}

func main() {
	testClac()
}
