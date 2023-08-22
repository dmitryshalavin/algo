package main

import "fmt"

func bubbleSort(a []int) {
	if len(a) < 2 {
		return
	}

	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func main() {
	a := []int{5, 1, 2, 7, 8, 3, 1}
	bubbleSort(a)
	fmt.Println(a)
}
