package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	dp := make([]int, len(nums))

	for i := 1; i <= nums[0] && i < len(nums); i++ {
		dp[i] = 1
	}

	// fmt.Println(dp)
	for i := 1; i < len(nums); i++ {
		for j := i; j <= nums[i]+i && j < len(nums); j++ {
			if dp[j] == 0 {
				dp[j] = dp[i] + 1
			}
			if dp[len(dp)-1] != 0 {
				return dp[len(dp)-1]
			}
		}
		// fmt.Println(dp)
	}

	return dp[len(dp)-1]
}
