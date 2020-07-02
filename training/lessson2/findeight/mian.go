package main

import "fmt"

func sumEight(a [8]int) [][]int {
	r := [][]int{}
	for i, _ := range a {
		for j := i + 1; j < len(a); j++ {
			if j == i {
				continue
			}
			if a[j]+a[i] == 8 {
				r = append(r, []int{i, j})
			}
		}
	}
	return r
}

func main() {
	a := [8]int{0, 1, 3, 5, 7, 8, 2, 6}
	fmt.Println(sumEight(a))
}
