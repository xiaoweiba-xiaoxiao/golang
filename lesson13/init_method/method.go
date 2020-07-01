package main

import "fmt"

type People struct {
	Name    string
	Country string
}

func (p People) Print() {
	fmt.Printf("name=%s country=%s\n", p.Name, p.Country)
}

func (p *People) set(name, country string) {
	p.Name = name
	p.Country = country
}

func main() {
	var p1 People = People{
		Name:    "people01",
		Country: "China",
	}
	p1.Print()
	p1.set("people02", "English")
	p1.Print()
}
