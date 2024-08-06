package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calculate("1+2*5/3+6/4*2"))
}

func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	var (
		stack = make([]string, 0, len(s))
		next  string
	)

	result, start := getNextNum(s, 0)
	stack = append(stack, result)
	for i := start + 1; i < len(s); i++ {
		ind := i
		next, i = getNextNum(s, i+1)

		switch s[ind] {
		case '+':
			stack = append(stack, "+")
			stack = append(stack, next)
		case '-':
			stack = append(stack, "-")
			stack = append(stack, next)
		case '*':
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			val1, _ := strconv.Atoi(last)
			val2, _ := strconv.Atoi(next)

			stack = append(stack, fmt.Sprint(val1*val2))
		case '/':
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			val1, _ := strconv.Atoi(last)
			val2, _ := strconv.Atoi(next)

			stack = append(stack, fmt.Sprint(val1/val2))
		}
	}

	val, _ := strconv.Atoi(stack[0])
	stack = stack[1:]

	for len(stack) > 1 {
		val2, _ := strconv.Atoi(stack[1])
		op := stack[0]
		stack = stack[2:]
		if op == "+" {
			val += val2
		} else {
			val -= val2
		}
	}

	return val
}

func getNextNum(s string, pos int) (string, int) {
	var (
		num string
		ind int
	)
	for i := pos; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num += string(s[i])
			ind = i
			continue
		}
		break
	}

	return num, ind
}
