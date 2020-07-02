package main

import "fmt"

func selectSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func main() {
	i := []int{8, 3, 0, 2, 7, 10, 9, 4, 5}
	fmt.Println(i)
	fmt.Println(selectSort(i))
}
