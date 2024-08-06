package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello world")
}

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	var (
		stack = make([]string, 0, len(s))
		brack = make([]int, 0, len(s)/4)
		num   strings.Builder
	)
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if num.Len() != 0 {
				stack = append(stack, num.String())
			}
			last := brack[len(brack)-1]
			result := calculateFromStack(stack[last+1:])
			brack = brack[:len(brack)-1]
			stack = stack[:last]
			stack = append(stack, fmt.Sprint(result))
			num.Reset()
		} else if s[i] == '(' {
			if num.Len() != 0 {
				stack = append(stack, num.String())
			}
			stack = append(stack, string(s[i]))
			brack = append(brack, len(stack)-1)
			num.Reset()
		} else if s[i] == '+' || s[i] == '-' {
			if num.Len() != 0 {
				stack = append(stack, num.String())
			}
			num.Reset()
			stack = append(stack, string(s[i]))
		} else {
			num.WriteByte(s[i])
		}
	}

	stack = append(stack, num.String())

	return calculateFromStack(stack)
}

func calculateFromStack(stack []string) int {
	var result int
	for len(stack) > 0 {
		if stack[0] == "-" {
			val, _ := strconv.Atoi(stack[1])
			result -= val
			stack = stack[2:]
		} else if stack[0] == "+" {
			val, _ := strconv.Atoi(stack[1])
			result += val
			stack = stack[2:]
		} else {
			val, _ := strconv.Atoi(stack[0])
			stack = stack[1:]
			result += val
		}
	}

	return result
}
