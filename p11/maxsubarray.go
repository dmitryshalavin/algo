package main

import "fmt"

func MaxSubArray(a []int) ([]int, int) {
	dp := []int{}
	dpS := 0

	if len(a) == 0 {
		return dp, dpS
	}

	maxA, max := 0, 0
	left, right, idx := 0, 0, 0

	for i := 0; i < len(a); i++ {
		max += a[i]
		if max < a[i] {
			max = a[i]
			idx = i
		}

		if max > maxA {
			maxA = max
			left, right = idx, i
		}
	}

	return a[left : right+1], maxA
}

func main() {
	fmt.Println(MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
