package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minDifference([]int{1, 5, 0, 10, 14}))
}

func minDifference(nums []int) int {
	if len(nums) <= 3 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	result := math.MaxInt32
	result = min(result, nums[0]-nums[len(nums)-4])
	result = min(result, nums[1]-nums[len(nums)-3])
	result = min(result, nums[2]-nums[len(nums)-2])
	result = min(result, nums[3]-nums[len(nums)-1])

	return result
}
