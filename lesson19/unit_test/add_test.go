package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var a int = 10
	var b int = 20
	c := Add(a, b)
	if c != 30 {
		t.Fatalf("%d + %d = %d", a, b, c)
	}
	t.Logf("%d", c)
}

func TestSub(t *testing.T) {
	var a int = 20
	var b int = 10
	c := Sub(a, b)
	if c != 10 {
		t.Fatalf("%d - %d = %d", a, b, c)
	}
	t.Logf("%d", c)
}
