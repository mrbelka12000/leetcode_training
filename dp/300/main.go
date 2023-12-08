package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	_, max := incSeqSolution(nums)
	return max
}

func incSeqSolution(nums []int) ([]int, int) {
	dp := make([]int, len(nums))
	dp[0] = 1

	var result int = 1
	for i := 1; i < len(nums); i++ {
		maxInt := 1
		var check bool
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				maxInt = max(dp[j], maxInt)
				check = true
			}
		}
		if check {
			maxInt++
		}
		dp[i] = maxInt
		result = max(result, maxInt)
	}
	// fmt.Println(dp)

	return nil, result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
