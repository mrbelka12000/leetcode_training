package main

import "fmt"

func main() {
	fmt.Println(totalSteps([]int{7, 14, 4, 14, 13, 2, 6, 13}))
}

func totalSteps(nums []int) int {
	var result int

	var (
		stack []int
		dp    = make([]int, len(nums))
	)

	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			dp[i] = max(dp[stack[len(stack)-1]], dp[i]+1)
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	// fmt.Println(dp)
	for _, v := range dp {
		result = max(result, v)
	}

	return result
}
