package clac

import "fmt"

type TestUser struct {
	Name    string
	Sex     string
	Adrress string
	Age     int
	myself  string //filed 如果是小写则包外不能访问
}

//自定义构造函数
func New(name, sex, adrress string, age int) *TestUser {
	var testuser TestUser = TestUser{
		Name:    name,
		Sex:     sex,
		Adrress: adrress,
		Age:     age,
	}
	return &testuser
}

func (t *TestUser) Say() string {
	saystring := fmt.Sprintf("hello ervery one,my name is %s,I am a %s, I lived on %s,I am %d year old", t.Name, t.Sex, t.Adrress, t.Age)
	return saystring

}
