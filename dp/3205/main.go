package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxScore([]int{4, 5, 2, 8, 9, 1, 3}))
}

func maxScore(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j >= i; j-- {
			dp[j] = max(dp[j], dp[i]+(j-i)*nums[j])
		}
	}

	return dp[len(nums)-1]
}
