package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseParentheses("(ed(et(oc))el)"))
}

func reverseParentheses(s string) string {
	b := []byte(s)
	var stack []int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		}
		if s[i] == ')' {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			ind := i
			for last < ind {
				b[last], b[ind] = b[ind], b[last]
				last++
				ind--
			}
		}
	}

	result := strings.Builder{}
	for _, v := range b {
		if v == '(' || v == ')' {
			continue
		}
		result.WriteByte(v)
	}

	return result.String()
}
