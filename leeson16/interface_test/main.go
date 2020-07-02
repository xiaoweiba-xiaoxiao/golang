package main

import (
	"fmt"
)

type Xiannv interface {
	Butiful()
	Lovely()
}

type xiaoxiao struct {
}

func (x *xiaoxiao) Butiful() {
	fmt.Println("xiaoxiao is butiful girl")
}

func (x *xiaoxiao) Lovely() {
	fmt.Println("xiaoxiao is lovely girl")
}

func main() {
	var a xiaoxiao
	var b Xiannv
	b = &a
	b.Butiful()
	b.Lovely()
}
