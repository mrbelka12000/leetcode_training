package main

import (
	"fmt"
)

func main() {
	fmt.Println(maximumDifference([]int{1, 5, 2, 10}))
}

func maximumDifference(nums []int) int {
	var result int = -1
	minVal := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-minVal != 0 {
			result = max(result, nums[i]-minVal)
		}
		minVal = min(minVal, nums[i])
	}

	return result
}
