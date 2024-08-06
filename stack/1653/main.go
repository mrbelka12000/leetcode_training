package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumDeletions("aababbab"))
}

func minimumDeletions(s string) int {
	var (
		stack  []byte
		result int
	)
	for i := len(s) - 1; i >= 0; i-- {
		if len(stack) != 0 && s[i] == 'b' && stack[len(stack)-1] == 'a' {
			result++
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, s[i])
	}

	return result
}
