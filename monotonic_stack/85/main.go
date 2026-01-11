package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximalRectangle([][]byte{{'0', '1', '1'}}))
}

func maximalRectangle(matrix [][]byte) int {
	heights := make([]int, len(matrix[0]))
	var maxArea int
	for _, row := range matrix {
		for j, v := range row {
			if v == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}

		maxArea = max(maxArea, largetsRect(len(matrix[0]), heights))
	}
	return maxArea
}

func largetsRect(n int, heights []int) int {

	left := make([]int, n)
	for i := range left {
		left[i] = -1
	}
	var stack []int

	for i, h := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	right := make([]int, n)
	for i := 0; i < len(right); i++ {
		right[i] = n
	}
	stack = nil

	for i := n - 1; i >= 0; i-- {
		h := heights[i]

		for len(stack) > 0 && heights[stack[len(stack)-1]] >= h {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	var maxArea int
	for i, h := range heights {
		maxArea = max(maxArea, h*(right[i]-left[i]-1))
	}
	return maxArea
}
