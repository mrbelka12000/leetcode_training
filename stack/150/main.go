package main

import (
	"fmt"
	"strconv"
)

func main() {

	//fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))

	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))

}

var mp = map[string]struct{}{
	"*": {},
	"/": {},
	"+": {},
	"-": {},
}

func evalRPN(tokens []string) int {
	var result int

	var stack []int

	for i := 0; i < len(tokens); i++ {
		if _, ok := mp[tokens[i]]; ok {
			result = do(stack[len(stack)-2], stack[len(stack)-1], tokens[i])
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = result
			continue
		}

		val, _ := strconv.Atoi(tokens[i])
		stack = append(stack, val)
	}
	if len(tokens) == 1 {
		return stack[0]
	}

	return result
}

func do(a, b int, s string) int {
	switch s {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		fmt.Println(a / b)
		return a / b
	}

	return 0
}
