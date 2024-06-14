package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}

func findBuildings(heights []int) []int {
	if len(heights) == 1 {
		return []int{0}
	}
	var (
		stack  []int
		result []int
	)
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] <= heights[i] {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	for _, v := range stack {
		result = append(result, v)
	}

	return result
}
