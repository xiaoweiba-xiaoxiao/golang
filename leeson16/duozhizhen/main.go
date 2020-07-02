package main

import (
	"fmt"
)

type animo interface {
	Eat()
	Talk()
	Sleep()
}

type Describe interface {
	Describe()
}

type puru interface {
	animo
	Describe
}

type sheep struct {
	name string
}

func (s *sheep) Eat() {
	fmt.Println("I am eating")
}

func (s *sheep) Talk() {
	fmt.Println("I am sheep")
}

func (s *sheep) Sleep() {
	fmt.Println("I am sleeping")
}

func (s *sheep) Describe() {
	fmt.Println("sheep is puru")
}

func main() {
	var s *sheep
	var a puru
	a = s
	a.Eat()
	a.Talk()
	a.Sleep()
	a.Describe()
}
