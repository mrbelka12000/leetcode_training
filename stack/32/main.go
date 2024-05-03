package main

import "fmt"

func main() {
	fmt.Println(longestValidParentheses(")()())"))

}

func longestValidParentheses(s string) int {
	var stack []int
	tmp := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if len(stack) != 0 {
			if s[stack[len(stack)-1]] != '(' {
				continue
			} else {
				tmp[i] = 1
				tmp[stack[len(stack)-1]] = 1
				stack = stack[:len(stack)-1]
			}
		}
	}

	var (
		result int
		count  int
	)

	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 1 {
			count++
		} else {
			result = max(result, count)
			count = 0
		}
	}

	return max(result, count)
}
