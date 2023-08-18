package main

import "fmt"

func Check(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 == 1 {
		return false
	}

	stack := []rune{}

	for _, c := range s {
		if c == '(' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] == ')' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(Check("(())()(("))
}
