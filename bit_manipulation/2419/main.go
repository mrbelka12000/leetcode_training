package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestSubarray([]int{1, 2, 3, 3, 2, 2}))
}

func longestSubarray(nums []int) int {
	var maxVal int
	for _, v := range nums {
		maxVal = max(maxVal, v)
	}
	var (
		result, ind int
	)
	ind--
	for i := 0; i < len(nums); i++ {
		if nums[i] == maxVal {
			result = max(result, i-ind)
		} else {
			ind = i
		}
	}

	return max(result, 1)
}
