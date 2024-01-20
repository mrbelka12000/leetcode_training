package main

import (
	"fmt"
)

func main() {
	//	qwedfsadfdsafdasfa 0  1  2  3  4  5  6  7  8  9  10 11
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}), "result")

	// 1234dafdsfdsafads   1  2  3  4  5  6
	//fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))

	//fmt.Println(trap([]int{4, 2, 3}))

	//fmt.Println(trap([]int{4, 2, 0, 3, 2, 4, 3, 4}))
}

/*

0		    0   0
0		0   0 0 0
0 0   0 0 0 0 0 0
0 0   0 0 0 0 0 0
*/

/*

	      0
0		  0
0     0   0
0 0   0 0 0
0 0   0 0 0
*/

func trap(height []int) int {
	var (
		result int
		stack  []int
		lefts  []int
		rights = make([]int, len(height))
	)

	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && stack[len(stack)-1] < height[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			if height[i] < stack[len(stack)-1] {
				lefts = append(lefts, stack[len(stack)-1])
				continue
			}
		}
		stack = append(stack, height[i])
		lefts = append(lefts, stack[len(stack)-1])
	}

	stack = make([]int, 0)

	for i := len(height) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < height[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			if height[i] < stack[len(stack)-1] {
				rights[i] = stack[len(stack)-1]
				continue
			}
		}
		stack = append(stack, height[i])
		rights[i] = stack[len(stack)-1]
	}
	// fmt.Println(lefts)
	// fmt.Println(rights)

	for i := 0; i < len(lefts); i++ {
		lefts[i] = min(lefts[i], rights[i])
	}

	for i := 0; i < len(height); i++ {
		result += lefts[i] - height[i]
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
