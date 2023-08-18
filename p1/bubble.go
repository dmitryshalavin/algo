package main

import "fmt"

func bubbleSort(a []int) {
	if len(a) < 2 {
		return
	}

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func main() {
	a := []int{5, 1, 2, 7, 8, 3}
	bubbleSort(a)
	fmt.Println(a)
}
