package main

import (
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var a int = 10
		var d int = 20
		c := Add(a, d)
		if c != 30 {
			b.Fatalf("%d + %d = %d is wrong", a, d, c)
		}
		b.Logf("%d", c)
	}
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := 10
		d := 20
		c := Sub(d, a)
		if c != 10 {
			b.Fatalf("%d - %d = %d is wrong", a, d, c)
		}
		b.Logf("%d", c)
	}
}
