package main

import "fmt"

func spiral(n int) {
	s := make([][]int, n)
	for i := range s {
		s[i] = make([]int, n)
	}

	sideLen := n - 1
	step := 0
	dx, dy := 1, 0
	i, j := 0, 0
	for k := 0; k < n*n; k++ {
		s[j][i] = k + 1
		step++
		i += dx
		j += dy
		if sideLen == step {
			step = 0
			if dx == 1 {
				dx, dy = 0, 1
			} else if dy == 1 {
				dx, dy = -1, 0
			} else if dx == -1 {
				dx, dy = 0, -1
			} else { // dy == -1
				i++
				j++
				dx, dy = 1, 0
				sideLen -= 2
			}
		}

	}

	for i := range s {
		fmt.Println(s[i])
	}
}

func main() {
	spiral(5)
}
