package main

import "fmt"

func SelectSort(a []int) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[i], a[min] = a[min], a[i]
	}
}

func main() {
	a := []int{5, 1, 2, 7, 8, 3}
	SelectSort(a)
	fmt.Println(a)
}
