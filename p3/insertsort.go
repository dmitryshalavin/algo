package main

import "fmt"

func InsertSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				x := a[i]
				copy(a[j+1:], a[j:i])
				a[j] = x
			}
		}
	}
}

func main() {
	a := []int{5, 1, 2, 7, 8, 3}
	InsertSort(a)
	fmt.Println(a)
}
