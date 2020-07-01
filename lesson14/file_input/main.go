package main

import (
	"fmt"
	"os"
)

func testFscanf() {
	var a int
	var b string
	var c float32
	fmt.Fscanf(os.Stdin, "%d%s%f", &a, &b, &c)
	fmt.Fprintf(os.Stdout, "a=%d b=%s c=%f\n", a, b, c)
}

func testFscan() {
	var a int
	var b string
	var c float32
	fmt.Fscan(os.Stdin, &a, &b, &c)
	fmt.Fprintf(os.Stdout, "a=%d b=%s c=%f\n", a, b, c)
}

func testFscanln() {
	var a int
	var b string
	var c float32
	fmt.Fscanln(os.Stdin, &a, &b, &c)
	fmt.Fprintf(os.Stdout, "a=%d b=%s c=%f\n", a, b, c)
}

func main() {
	// testFscanf()
	// testFscan()
	testFscanln()
}
