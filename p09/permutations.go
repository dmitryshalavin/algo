package main

import "fmt"

func helper(arr []int, start int, end int) {
	if start == end {
		fmt.Println(arr)
		return
	}
	for i := start; i < end; i++ {
		arr[i], arr[start] = arr[start], arr[i]
		helper(arr, start+1, end)
		arr[i], arr[start] = arr[start], arr[i]
	}
}

func Permitations(n int) {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	helper(arr, 0, n)
}
func main() {
	Permitations(3)
}
