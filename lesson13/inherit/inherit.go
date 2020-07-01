package main

import "fmt"

//继承
type Animal struct {
	Name string
	Sex  string
}

type Dog struct {
	feet string
	*Animal
}

func (a *Animal) Talk() {
	fmt.Printf("I am %s,I am %s\n", a.Name, a.Sex)
}

func (d *Dog) Eat() {
	fmt.Println("I am eat")
}
func testdog() {
	var d *Dog = &Dog{
		feet: "four feets",
		Animal: &Animal{
			Name: "汪汪",
			Sex:  "母",
		},
	}
	d.Eat()
	d.Talk()
}

//多继承与冲突
type Samoye struct {
	Kind string
	*Dog
	*Animal
}

func (a *Animal) Feet() {
	fmt.Println("animal feet")
}

func (d *Dog) Feet() {
	fmt.Println("dog feet")
}

/* 存在是方法是优先以此方法生效，不存在是直接调用Samoye.Feet()就会冲突
func (s *Samoye)Feet(){
	fmt.Println("samoye feet")
}
*/
func testsamoye() {
	var samoye *Samoye = &Samoye{
		Kind:   "微笑天使",
		Animal: &Animal{},
		Dog:    &Dog{},
	}
	/* 由于Samoye.Feet()方法不存在，所以冲突
	samoye.Feet()
	*/
	//通过filed对象调用Feet()方法解决冲突
	samoye.Animal.Feet()
	samoye.Dog.Feet()
}

func main() {
	// testdog()
	testsamoye()
}
