package main

import "fmt"

func main() {
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
}

func maxAscendingSum(nums []int) int {
	var result int

	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			result = max(result, curr)
			curr = 0
		}
		curr += nums[i]
	}

	return max(result, curr)
}
