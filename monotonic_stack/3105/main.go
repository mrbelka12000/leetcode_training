package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestMonotonicSubarray([]int{1, 4, 3, 3, 2}))
}

func longestMonotonicSubarray(nums []int) int {
	return max(getLongest(nums, false, 0), getLongest(nums, true, math.MaxInt32))
}

func getLongest(nums []int, cond bool, prev int) int {

	var (
		result, count int
	)

	for i := 0; i < len(nums); i++ {
		if cond {
			if nums[i] >= prev {
				prev = 0
				count = 0
			}
		} else {
			if nums[i] <= prev {
				prev = 0
				count = 0
			}
		}
		prev = nums[i]
		count++
		result = max(result, count)
	}

	return result
}
