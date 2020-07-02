package main

import "fmt"

func arraySum(a *[5]int) int {
	sum := 0
	for _, val := range a {
		sum += val
	}
	a[4] = 100
	return sum
}

func main() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("sum of a =%d\n", arraySum(&a))
	fmt.Println(a)
}
