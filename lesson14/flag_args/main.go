package main

import (
	"flag"
	"fmt"
)

var (
	level int
	test  bool
	name  string
)

func init() {
	flag.StringVar(&name, "s", "default", "the name")
	flag.IntVar(&level, "l", 1, "the level")
	flag.BoolVar(&test, "b", false, "the test")
	flag.Parse()
}

func main() {
	fmt.Printf("level=%d\n", level)
	fmt.Printf("test=%v\n", test)
	fmt.Printf("%s\n", name)
}
