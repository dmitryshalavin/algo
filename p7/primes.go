package main

import "fmt"

func Primes(n int) []int {
	a := make([]bool, n+2)

	for i := 2; i <= n; i++ {
		if a[i] {
			continue
		}
		for j := i * i; j <= n+1; j += i {
			a[j] = true
		}
	}

	res := []int{}
	for i, v := range a[2:] {
		if !v {
			res = append(res, i+2)
		}
	}

	return res
}

func main() {
	fmt.Println(Primes(99))
}
