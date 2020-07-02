package main

import "fmt"

func bubleSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}
func main() {
	i := []int{8, 3, 0, 2, 7, 10, 9, 4, 5}
	fmt.Println(i)
	fmt.Println(bubleSort(i))
}
