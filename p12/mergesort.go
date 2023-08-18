package main

import "fmt"

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	return helper(a[0:len(a)/2], a[len(a)/2:])
}

func helper(a []int, b []int) []int {
	lA, lB := len(a), len(b)
	if lA < 3 {
		if len(a) == 2 {
			if a[0] > a[1] {
				a[0], a[1] = a[1], a[0]
			}
		}
	} else {
		a = helper(a[0:lA/2], a[lA/2:])
	}

	if lB < 3 {
		if lB == 2 {
			if b[0] > b[1] {
				b[0], b[1] = b[1], b[0]
			}
		}
	} else {
		b = helper(b[0:lB/2], b[lB/2:])
	}

	c := make([]int, lA+lB)
	i, j := 0, 0
	for k := 0; k < len(c); k++ {
		if i < lA && j < lB {
			if a[i] < b[j] {
				c[k] = a[i]
				i++
			} else {
				c[k] = b[j]
				j++
			}
		} else if i < lA {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}

	}

	return c
}

func main() {
	fmt.Println(MergeSort([]int{2, 6, 1, 5, -3, 9, 7}))
}
