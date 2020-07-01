package main

import (
	"encoding/json"
	"fmt"

	"xiao.com/golang/lesson/lesson12/testtag"
)

type User struct {
	Username string
	Sex      string
	Age      int
}

//初始化struct方法1
func testUser1() {
	var user User
	user.Username = "xiaoxiao"
	user.Sex = "girl"
	user.Age = 29
	fmt.Printf("user: %#v", user)
}

//初始化struct方法2
func testUser2() {
	var user User = User{
		Username: "xiaoxiao",
		Sex:      "girl",
		Age:      29,
	}
	fmt.Printf("user: %#v", user)
}

//struct初始化的默认值
func testUser3() {
	var user User
	fmt.Printf("user: %#v", user)
}

//strcut指针
func testUser4() {
	/*
		var user *User //此时指针为空指针
		user = &User{} //指针一定要初始化
	*/
	var user *User = &User{}
	// 或user :=&User{}
	user.Username = "xiaoxiao"
	user.Sex = "girl"
	user.Age = 29
	fmt.Printf("user: %#v", user)
}

//strcut指针
func testUser5() {
	var user *User = &User{
		Username: "xiaoxiao",
		Sex:      "girl",
		Age:      29,
	}
	fmt.Printf("user: %#v", user)
}

//struct 指针
func testUser6() {
	var user *User = new(User)
	user.Username = "xiaochao"
	user.Sex = "boy"
	user.Age = 31
	fmt.Printf("user: %#v", user)
}

//stuct 字段是连续的指针地址
func testUser7() {
	var user *User = new(User)
	fmt.Printf("%p\n", &user.Username)
	fmt.Printf("%p\n", &user.Sex)
	fmt.Printf("%p\n", &user.Age)
}

//struct 构造函数请参照clac和testclac包

//struct 匿名字段
type User2 struct {
	Name string
	Age  int
	string
	//string 匿名字段同一种类型只允许有一个
	int
}

func testUser8() {
	user := User2{
		Name:   "xiao",
		Age:    28,
		string: "xiaochaoren",
		int:    28,
	}
	fmt.Printf("%#v,user string=%s,user int=%d", user, user.string, user.int)
}

//struct 嵌套

type Adress struct {
	Province  string
	City      string
	Creattime string
}

type User3 struct {
	Name string
	Age  int
	Addr Adress
}

func testUser9() {
	var user User3 = User3{
		Name: "xiao",
		Age:  28,
		Addr: Adress{
			Province: "hubei",
			City:     "wuhan",
		},
	}
	fmt.Printf("%#v\n", user)
}

//struct .匿匿名结构体 匿匿名结构体与继承
type User4 struct {
	Name string
	Sex  string
	Adress
}

func testUser10() {
	var user User4 = User4{
		Name: "xiaoxiao",
		Sex:  "girl",
		Adress: Adress{
			Province: "hubei",
			City:     "wuhan",
		},
	}
	fmt.Printf("%#v", user)
}

//冲突和冲突解决
type Email struct {
	Emailaddr string
	Creattime string
}

type User5 struct {
	Name string
	Sex  string
	Adress
	Email
	// Creattime string //此行注释就会引起冲突
}

func testUser11() {
	user := User5{
		Name: "xiaoxiao",
		Sex:  "girl",
		Adress: Adress{
			Province:  "hubei",
			City:      "wuhan",
			Creattime: "2020-07-02 18:00:00",
		},
		Email: Email{
			Emailaddr: "xiaoxiao@xiao.com",
			Creattime: "2020-07-02 19:00:00",
		},
		// Creattime: "2020-07-02 20:00:00",
	}
	fmt.Printf("%s", user.Email.Creattime) //解决冲突的办法
}

//tag 请参见testtag包
func testUser12() {
	var user *testtag.Testtag = testtag.New("xiaoxiao", "girl", 29)
	data, _ := json.Marshal(user)
	fmt.Printf("%s", string(data))

}

func main() {
	// testUser1()
	// testUser2()
	// testUser3()
	// testUser4()
	// testUser5()
	// testUser6()
	// testUser7()
	// testUser8()
	// testUser9()
	// testUser10()
	// testUser11()
	testUser12()
}
